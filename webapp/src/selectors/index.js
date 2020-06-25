import Constants from '../constants';

const getPluginState = (state) =>
    state[`plugins-${Constants.PLUGIN_NAME}`] || {};

const survey = (state) => getPluginState(state).survey;

export default {
    survey,
};
