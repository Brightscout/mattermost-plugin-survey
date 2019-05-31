import React from 'react';
import PropTypes from 'prop-types';

import {FormGroup, ControlLabel, FormControl} from 'react-bootstrap';
import {changeOpacity, makeStyleFromTheme} from 'mattermost-redux/utils/theme_utils';

import './styles.css';

export default class QuestionTypeLikertScale extends React.PureComponent {
    static propTypes = {
        id: PropTypes.string.isRequired,
        text: PropTypes.string.isRequired,
        theme: PropTypes.object.isRequired,
        index: PropTypes.number.isRequired,
        responses: PropTypes.array.isRequired,
        handleChange: PropTypes.func.isRequired,
    };

    constructor(props) {
        super(props);
        this.state = {
            selectedValue: '',
            hoveredValue: '',
        };
    }

    handleChange = (evt) => {
        const selectedValue = evt.target.value;
        this.props.handleChange(this.props.id, selectedValue);
        this.setState({
            selectedValue,
        });
    };

    handleMouseEnter = (e) => {
        this.setState({
            hoveredValue: e.target.dataset.value,
        });
    };

    handleMouseLeave = () => {
        this.setState({
            hoveredValue: '',
        });
    };

    render() {
        const {id: questionID, index, responses, text, theme} = this.props;
        const {hoveredValue, selectedValue} = this.state;
        const style = getStyle(theme);

        const radios = responses.map((response, idx) => {
            let optionStyle;
            if (selectedValue === response.value) {
                optionStyle = style.selected;
            } else if (hoveredValue === response.value) {
                optionStyle = style.hovered;
            }

            const labelID = questionID + response.value;
            const radioButtonID = `${idx}${questionID}`;
            return (
                <ControlLabel
                    id={labelID}
                    key={labelID}
                    className='likert-option'
                    style={optionStyle}
                    htmlFor={radioButtonID}
                    data-value={response.value}
                    onMouseOver={this.handleMouseEnter}
                    onMouseOut={this.handleMouseLeave}
                >
                    <FormControl
                        type={'radio'}
                        aria-labelledby={labelID}
                        name={questionID}
                        id={radioButtonID}
                        className='display-none'
                        value={response.value}
                        onClick={this.handleChange}
                    />
                    <span className='likert-option-label'>{response.text}</span>
                </ControlLabel>
            );
        });

        return (
            <FormGroup>
                <p id={questionID}>{`${index}. ${text}`}</p>
                <div
                    className='likert-options'
                    aria-labelledby={questionID}
                >
                    {radios}
                </div>
            </FormGroup>
        );
    }
}

const getStyle = makeStyleFromTheme((theme) => ({
    selected: {
        backgroundColor: theme.sidebarTextActiveBorder,
        color: theme.sidebarTextActiveColor,
    },
    hovered: {
        backgroundColor: changeOpacity(theme.sidebarTextHoverBg, 0.1),
    },
}));