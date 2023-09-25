import Vuex from 'vuex';

import * as openapi from '@/openapi';

import API from '@/api';

export default {
    actions: {
        async loadAlerts(ctx: any) {
            try {
                const resp = await API.getAlertList();

                ctx.commit('saveAlerts', resp.data.items);
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        },

        async fixAlert(ctx: any, id: string) {
            try {
                await API.fixAlert(id);

                ctx.commit('fixAlert', id);
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        }
    },

    mutations: {
        saveAlerts(state: any, alerts: Array<openapi.GetAlertListResponseItemsInner>) {
            state.alerts = alerts;
        },

        fixAlert(state: any, id: string) {
            state.alerts = state.alerts.filter((value: openapi.GetAlertListResponseItemsInner) => value.uuid != id)
        }
    },

    state: {
        alerts: Array<openapi.GetAlertListResponseItemsInner>(),
    },

    getters: {
        alerts(state: any) {
            return state.alerts
        },
    },
};