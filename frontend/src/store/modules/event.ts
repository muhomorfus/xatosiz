import Vuex from 'vuex';

import * as openapi from '@/openapi';

import API from '@/api';

export default {
    actions: {
        async loadEvent(ctx: any, uuid: string) {
            try {
                const resp = await API.getEvent(uuid);

                ctx.commit('saveEvent', resp.data);
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        },

        async fix(ctx: any) {
            try {
                await API.fixEvent(ctx.state.event.uuid);

                ctx.commit('fixEvent');
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        }
    },

    mutations: {
        saveEvent(state: any, event: openapi.Event) {
            state.event = event;
        },

        fixEvent(state: any) {
            state.event.fixed = true;
        }
    },

    state: {
        event: {} as openapi.Event,
    },

    getters: {
        event(state: any) {
            return state.event;
        }
    },
};