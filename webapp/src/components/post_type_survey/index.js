import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {getCurrentUser} from 'mattermost-redux/selectors/entities/common';
import {
    getCurrentTeamId,
    getTeam,
} from 'mattermost-redux/selectors/entities/teams';

import Actions from '../../actions';
import Client from '../../client';
import Selectors from '../../selectors';

import PostTypeSurvey from './post_type_survey';

const mapStateToProps = (state) => {
    const currentTeam = getTeam(state, getCurrentTeamId(state));
    return {
        currentUser: getCurrentUser(state) || {},
        dashboardURL: Client.getDashboardURL(
            currentTeam.name,
            Selectors.dashboardPath(state)
        ),
    };
};

const mapDispatchToProps = (dispatch) => bindActionCreators({
    openRiffDashboard: Actions.openRiffDashboard,
    openSurveyModal: Actions.openSurveyModal,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(PostTypeSurvey);
