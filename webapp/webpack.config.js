const exec = require('child_process').exec;

const path = require('path');

module.exports = {
    entry: [
        './src/index.jsx',
    ],
    resolve: {
        alias: {
            assets: path.resolve(__dirname, '../assets'),
            constants: path.resolve(__dirname, 'src/constants'),
        },
        modules: [
            path.resolve(__dirname, 'src'),
            'node_modules',
        ],
        extensions: ['*', '.js', '.jsx', '.json', '.scss'],
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        cacheDirectory: true,
                        plugins: [
                            '@babel/plugin-proposal-class-properties',
                            '@babel/plugin-syntax-dynamic-import',
                            '@babel/proposal-object-rest-spread',
                        ],
                        presets: [ // Babel configuration is in .babelrc because jest requires it to be there.
                            ['@babel/preset-env', {
                                targets: {
                                    chrome: 66,
                                    firefox: 60,
                                    edge: 42,
                                    safari: 12,
                                },
                                modules: false,
                                debug: false,
                                corejs: '3.6.4',
                                useBuiltIns: 'usage',
                                shippedProposals: true,
                            }],
                            ['@babel/preset-react', {
                                useBuiltIns: true,
                            }],
                        ],
                    },
                },
            },
            {
                test: /.(bmp|gif|jpe?g|png|svg)$/,
                exclude: /node_modules/,
                use: [
                    {
                        loader: 'file-loader',
                        options: {

                            // For assets, only emit the file URL
                            // as the static file are served from the plugin server.
                            emitFile: false,
                            name: '[name].[ext]',
                            publicPath: '/plugins/ai.riffanalytics.survey/static/',
                        },
                    },
                ],
            },
            {
                test: /\.css$/,
                use: [
                    'style-loader',
                    'css-loader',
                ],
            },
            {
                test: /\.scss$/,
                use: [
                    'style-loader',
                    'css-loader',
                    {
                        loader: 'sass-loader',
                        options: {
                            includePaths: ['node_modules/compass-mixins/lib', 'sass'],
                        },
                    },
                ],
            },
        ],
    },
    externals: {
        react: 'React',
        redux: 'Redux',
        'prop-types': 'PropTypes',
        'post-utils': 'PostUtils',
        'react-bootstrap': 'ReactBootstrap',
        'react-dom': 'ReactDOM',
        'react-redux': 'ReactRedux',
    },
    output: {
        path: path.join(__dirname, '/dist'),
        publicPath: '/',
        filename: 'main.js',
    },
    devtool: 'source-map',
    performance: {
        hints: 'warning',
    },
    target: 'web',
    plugins: [
        {
            apply: (compiler) => {
                compiler.hooks.afterEmit.tap('AfterEmitPlugin', () => {
                    exec('cd .. && make reset', (err, stdout, stderr) => {
                        if (stdout) {
                            process.stdout.write(stdout);
                        }
                        if (stderr) {
                            process.stderr.write(stderr);
                        }
                    });
                });
            },
        },
    ],
};
