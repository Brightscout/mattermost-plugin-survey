import React from 'react';
import PropTypes from 'prop-types';

import {formatText, messageHtmlToComponent} from 'post-utils';

import {Button} from 'react-bootstrap';

import './styles.css';

export default class PostTypeSurvey extends React.PureComponent {
    static propTypes = {
        dashboardURL: PropTypes.string.isRequired,
        post: PropTypes.object.isRequired,
        currentUser: PropTypes.object.isRequired,
        setCurrentPostID: PropTypes.func.isRequired,
        setCurrentPostProps: PropTypes.func.isRequired,
        openSurveyModal: PropTypes.func.isRequired,
        openRiffDashboard: PropTypes.func.isRequired,
    };

    constructor(props) {
        super(props);
        this.state = {
        };
    }

    openModal = () => {
        this.props.setCurrentPostID(this.props.post.id);
        this.props.setCurrentPostProps(this.props.post.props);
        this.props.openSurveyModal();
    };

    goToDashboard = () => {
        // window.location.href = this.props.dashboardURL;
        this.props.openRiffDashboard(this.props.post.props.meeting_id);
    };

    renderSubmitted = () => {
        const message = `Thanks for your feedback @${this.props.currentUser.username}! Have you checked out your Riff Stats in the Dashboard?`;
        return (
            <div>
                {messageHtmlToComponent(formatText(message, {atMentions: true}))}
                <Button
                    bsStyle='primary'
                    className='survey-action-button'
                    onClick={this.goToDashboard}
                >
                    {'View Dashboard'}
                </Button>
            </div>
        );
    };

    renderNotSubmitted = () => {
        const message = `Hi @${this.props.currentUser.username} - Please tell us about the meeting you just had. It only takes 30 seconds, and helps improve Riff.`;

        return (
            <div>
                {messageHtmlToComponent(formatText(message, {atMentions: true}))}
                <Button
                    bsStyle='primary'
                    aria-haspopup='true'
                    className='survey-action-button'
                    onClick={this.openModal}
                >
                    {'Start Survey'}
                </Button>
            </div>
        );
    };

    render() {
        const {post} = this.props;
        const postProps = post.props;
        if (postProps.submitted) {
            return this.renderSubmitted();
        }

        return this.renderNotSubmitted();
    }
}
