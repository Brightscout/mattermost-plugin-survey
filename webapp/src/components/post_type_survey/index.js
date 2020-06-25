import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {getCurrentUser} from 'mattermost-redux/selectors/entities/common';

import Actions from '../../actions';

import PostTypeSurvey from './post_type_survey';

const mapStateToProps = (state) => ({
    currentUser: getCurrentUser(state) || {},
});

const mapDispatchToProps = (dispatch) => bindActionCreators({
    openRiffDashboard: Actions.openRiffDashboard,
    openSurveyModal: Actions.openSurveyModal,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(PostTypeSurvey);
