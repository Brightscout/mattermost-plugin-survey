import React from 'react';
import PropTypes from 'prop-types';

import {
    Alert,
    Button,
    ButtonGroup,
    Clearfix,
    Modal,
    OverlayTrigger,
    Tooltip,
} from 'react-bootstrap';

import QuestionTypeOpen from '../question_type_open';
import QuestionTypeLikertScale from '../question_type_likert_scale';

import constants from '../../constants';
import loadingGif from '../../../assets/load.gif';

import './styles.css';

export default class SurveyModal extends React.PureComponent {
    static propTypes = {
        theme: PropTypes.object.isRequired,
        surveyOptions: PropTypes.object.isRequired,
        visible: PropTypes.bool.isRequired,
        close: PropTypes.func.isRequired,
        getSurvey: PropTypes.func.isRequired,
        submitSurveyResponses: PropTypes.func.isRequired,
    };

    constructor(props) {
        super(props);
        this.state = {
            survey: {
                id: '',
                version: '',
                title: '',
                description: '',
                questions: [],
            },
            responses: {},
            loading: false,
            loadingSubmit: false,
            getSurveyError: false,
            submitResponseError: false,
            validResponses: false,
        };
        this.submitErrorRef = React.createRef();
    }

    componentDidMount() {
        if (this.props.visible) {
            this.getSurvey();
        }
    }

    componentDidUpdate(prevProps, prevState) {
        if (this.props.visible && !prevProps.visible) {
            this.getSurvey();
        }

        if (this.state.submitResponseError && !prevState.submitResponseError) {
            // Scroll to the error alert.
            if (this.submitErrorRef.current) {
                this.submitErrorRef.current.scrollIntoView({
                    behavior: 'smooth',
                });
            }
        }
    }

    getSurvey = async () => {
        const {surveyOptions, getSurvey} = this.props;
        this.setState({
            loading: true,
            getSurveyError: false,
            validResponses: false,
            submitResponseError: false,
        });

        const {data} = await getSurvey(
            surveyOptions.surveyID,
            surveyOptions.surveyVersion,
            surveyOptions.meetingID,
        );
        if (data) {
            const survey = data;
            const responses = survey.questions.reduce((obj, question) => {
                obj[question.id] = '';
                return obj;
            }, {});

            this.setState({
                survey,
                responses,
                loading: false,
                validResponses: this.validateResponses(responses),
            });
        } else {
            this.setState({
                loading: false,
                getSurveyError: true,
            });
        }
    };

    handleClose = () => {
        this.props.close();
    };

    handleSubmit = async () => {
        const {survey, responses} = this.state;
        const {surveyOptions} = this.props;

        this.setState({
            loadingSubmit: true,
            submitResponseError: false,
        });
        const {data} = await this.props.submitSurveyResponses(
            surveyOptions.postID,
            surveyOptions.meetingID,
            survey.id,
            survey.version,
            responses
        );
        if (data) {
            this.setState({
                loadingSubmit: false,
            });
            this.handleClose();
        } else {
            this.setState({
                loadingSubmit: false,
                submitResponseError: true,
            });
        }
    };

    validateResponses = (responses) => {
        for (const key in responses) {
            if (responses.hasOwnProperty(key) && responses[key].trim() !== '') {
                return true;
            }
        }

        return false;
    };

    handleUpdateQuestionResponse = (questionID, response) => {
        this.setState((prevState) => {
            const responses = {...prevState.responses};
            responses[questionID] = response;
            return {
                validResponses: this.validateResponses(responses),
                submitResponseError: false,
                responses,
            };
        });
    };

    renderQuestions = () => {
        const {theme} = this.props;
        const questionsList = this.state.survey.questions;

        return questionsList.map((question, idx) => {
            const baseProps = {
                index: idx + 1,
                id: question.id,
                key: question.id,
                text: question.text,
                theme,
                handleChange: this.handleUpdateQuestionResponse,
            };
            switch (question.type) {
            case constants.QUESTION_TYPES.OPEN:
                return <QuestionTypeOpen {...baseProps}/>;

            case constants.QUESTION_TYPES.FIVE_POINT_LIKERT_SCALE:
                return (
                    <QuestionTypeLikertScale
                        {...baseProps}
                        responses={
                            constants.FIVE_POINT_LIKERT_SCALE_RESPONSES
                        }
                    />
                );

            default:
                return null;
            }
        });
    };

    renderLoading = () => {
        return (
            <div className='survey-loading'>
                <img
                    alt={'Loading'}
                    src={loadingGif}
                />
            </div>
        );
    };

    renderSubmitButton = () => {
        const {loadingSubmit, validResponses} = this.state;
        const disabled = loadingSubmit || !validResponses;

        let submitLoader;
        if (loadingSubmit) {
            submitLoader = (
                <span
                    className='fa fa-spinner fa-fw fa-pulse spinner'
                    title={'Loading Icon'}
                />
            );
        }

        const submitButton = (
            <Button
                type='submit'
                bsStyle='primary'
                className='submit-survey-btn'
                onClick={this.handleSubmit}
                disabled={disabled}
                style={disabled ? {pointerEvents: 'none'} : {}}
            >
                {submitLoader}
                {'Submit'}
            </Button>
        );

        if (!validResponses) {
            return (
                <OverlayTrigger
                    rootClose={true}
                    trigger={['hover', 'focus']}
                    placement='top'
                    overlay={
                        <Tooltip>
                            {constants.ERROR_MESSAGES.VALIDATE_SURVEY}
                        </Tooltip>
                    }
                >
                    <div className='survey-submit-button-container'>
                        {submitButton}
                    </div>
                </OverlayTrigger>
            );
        }

        return submitButton;
    };

    renderSurvey = () => {
        const {survey, loadingSubmit, submitResponseError} = this.state;
        const submitButton = this.renderSubmitButton();

        let errorAlert;
        if (submitResponseError) {
            errorAlert = (
                <React.Fragment>
                    <Alert
                        bsStyle='warning'
                        className='survey-submit-server-error-alert'
                    >
                        <i
                            className='fa fa-warning'
                            title='Server Error'
                        />
                        {constants.ERROR_MESSAGES.SUBMIT_SURVEY}
                    </Alert>
                    <div ref={this.submitErrorRef}/>
                </React.Fragment>
            );
        }

        const questions = this.renderQuestions();
        return (
            <div>
                <p className='survey-banner-text'>{survey.description}</p>
                {questions}
                <Clearfix>
                    <ButtonGroup className='float-right survey-modal-buttons'>
                        <Button
                            type='button'
                            bsStyle='secondary'
                            onClick={this.handleClose}
                            disabled={loadingSubmit}
                        >
                            {'Cancel'}
                        </Button>
                        {submitButton}
                    </ButtonGroup>
                </Clearfix>
                {errorAlert}
            </div>
        );
    };

    renderGetSurveyError = () => {
        return (
            <div className='survey-fetch-server-error'>
                <i
                    className='fa fa-warning'
                    title='Server Error'
                />
                {constants.ERROR_MESSAGES.GET_SURVEY}
            </div>
        );
    };

    renderCancelFooter = () => {
        return (
            <Modal.Footer>
                <Button
                    type='button'
                    bsStyle='secondary'
                    onClick={this.handleClose}
                >
                    {'Cancel'}
                </Button>
            </Modal.Footer>
        );
    };

    render() {
        const {survey, loading, getSurveyError} = this.state;
        const {visible} = this.props;

        let content;
        let cancelFooter;

        if (loading) {
            content = this.renderLoading();
        } else if (getSurveyError) {
            content = this.renderGetSurveyError();
            cancelFooter = this.renderCancelFooter();
        } else {
            content = this.renderSurvey();
        }

        return (
            <Modal
                aria-hidden={!visible}
                aria-labelledby='survey-modal-title'
                show={visible}
                onHide={this.handleClose}
                backdrop={'static'}
            >
                <Modal.Header
                    closeButton={true}
                    closeLabel={'Close'}
                >
                    <Modal.Title id='survey-modal-title'>
                        {survey.title}
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body className='survey-modal-body'>{content}</Modal.Body>
                {cancelFooter}
            </Modal>
        );
    }
}
