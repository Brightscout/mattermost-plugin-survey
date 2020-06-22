import Client from '../client';
import Constants from '../constants';

export const openRiffDashboard = (meetingID) => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.OPEN_RIFF_DASHBOARD,
        meetingID,
    });
};

// TODO: Remove this.
export function getDashboardPath() {
    return async (dispatch) => {
        let data;
        try {
            data = await Client.getDashboardPath();
        } catch (error) {
            return {
                data: null,
                error,
            };
        }

        dispatch({
            type: Constants.ACTION_TYPES.RECEIVED_DASHBOARD_PATH,
            data,
        });
        return {
            data,
            error: null,
        };
    };
}
