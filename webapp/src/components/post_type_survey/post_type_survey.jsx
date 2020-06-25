import React from 'react';
import PropTypes from 'prop-types';

import {formatText, messageHtmlToComponent} from 'post-utils';

import {Button} from 'react-bootstrap';

import './styles.css';

export default class PostTypeSurvey extends React.PureComponent {
    static propTypes = {
        post: PropTypes.object.isRequired,
        currentUser: PropTypes.object.isRequired,
        openSurveyModal: PropTypes.func.isRequired,
        openRiffDashboard: PropTypes.func.isRequired,
    };

    constructor(props) {
        super(props);
        this.state = {};
    }

    openModal = () => {
        const postID = this.props.post.id;
        const postProps = this.props.post.props;
        const meetingID = postProps.meeting_id;

        this.props.openSurveyModal(
            postID,
            meetingID,
            postProps.survey_id,
            postProps.survey_version
        );
    };

    goToDashboard = () => {
        this.props.openRiffDashboard(this.props.post.props.meeting_id);
    };

    renderSubmitted = () => {
        const message = `Thanks for your feedback @${this.props.currentUser.username}! Have you checked out your Riff Stats in the Dashboard?`;
        return (
            <div>
                {messageHtmlToComponent(
                    formatText(message, {atMentions: true})
                )}
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
                {messageHtmlToComponent(
                    formatText(message, {atMentions: true})
                )}
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
