import Vuex from 'vuex';

import * as openapi from '@/openapi';

import API from '@/api';

export default {
    actions: {
        async loadEvents(ctx: any) {
            try {
                const resp = await API.getEventList();

                ctx.commit('saveEvents', resp.data.items);
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        },
    },

    mutations: {
        saveEvents(state: any, events: Array<openapi.Event>) {
            state.events = events;
        }
    },

    state: {
        events: Array<openapi.Event>(),
    },

    getters: {
        allEvents(state: any) {
            return state.events;
        }
    },
};