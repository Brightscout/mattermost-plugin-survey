import request from 'superagent';
import Cookies from 'js-cookie';

import {buildQueryString} from 'mattermost-redux/utils/helpers';

import Constants from '../constants';

/**
 *  Add web utilities for interacting with servers here
 */
export default class Client {
    constructor() {
        const url = new URL(window.location.href);
        this.baseUrl = `${url.protocol}//${url.host}`;
        this.pluginUrl = `${this.baseUrl}/plugins/${Constants.PLUGIN_NAME}`;
        this.apiUrl = `${this.baseUrl}/api/v4`;
        this.pluginApiUrl = `${this.pluginUrl}/api/v1`;
    }

    getSurvey = async (surveyID, surveyVersion, meetingID) => {
        const queryParams = {
            survey_id: surveyID,
            survey_version: surveyVersion,
            meeting_id: meetingID,
        };
        const url = `${this.pluginApiUrl}/survey${buildQueryString(queryParams)}`;
        return this.doGet(url);
    };

    getSurveyResponses = async (meetingID) => {
        const url = `${this.pluginApiUrl}/meetings/${meetingID}/response`;
        return this.doGet(url);
    };

    submitSurveyResponses = (surveyPostID, meetingID, surveyID, surveyVersion, responses) => {
        const queryParams = {
            survey_post_id: surveyPostID,
        };
        const url = `${this.pluginApiUrl}/submit${buildQueryString(queryParams)}`;
        const body = {
            meeting_id: meetingID,
            survey_id: surveyID,
            survey_version: surveyVersion,
            responses,
        };
        return this.doPost(url, body);
    };

    doGet = async (url, headers = {}) => {
        headers['X-Requested-With'] = 'XMLHttpRequest';

        const response = await request
            .get(url)
            .set(headers)
            .type('application/json')
            .accept('application/json');

        return response.body;
    };

    doPost = async (url, body, headers = {}) => {
        headers['X-Requested-With'] = 'XMLHttpRequest';
        headers['X-CSRF-Token'] = Cookies.get(Constants.MATTERMOST_CSRF_COOKIE);

        const response = await request
            .post(url)
            .send(body)
            .set(headers)
            .type('application/json')
            .accept('application/json');

        return response.body;
    };

    doPut = async (url, body, headers = {}) => {
        headers['X-Requested-With'] = 'XMLHttpRequest';
        headers['X-CSRF-Token'] = Cookies.get(Constants.MATTERMOST_CSRF_COOKIE);

        const response = await request
            .put(url)
            .send(body)
            .set(headers)
            .type('application/json')
            .accept('application/json');

        return response.body;
    }

    doDelete = async (url, body, headers = {}) => {
        headers['X-Requested-With'] = 'XMLHttpRequest';
        headers['X-CSRF-Token'] = Cookies.get(Constants.MATTERMOST_CSRF_COOKIE);

        const response = await request
            .delete(url)
            .send(body)
            .set(headers)
            .type('application/json')
            .accept('application/json');

        return response.body;
    };
}
