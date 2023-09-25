import Vuex from 'vuex';

import groups from './modules/groups';
import events from './modules/events';
import event from './modules/event';
import errors from './modules/errors';
import group from './modules/group';
import alert_configs from './modules/alert_configs';
import alerts from './modules/alerts';

export default new Vuex.Store({
    modules: {
        groups,
        events,
        event,
        errors,
        group,
        alert_configs,
        alerts
    },
});