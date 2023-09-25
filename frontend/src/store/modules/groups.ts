import Vuex from 'vuex';

import * as openapi from '@/openapi';

import API from '@/api';

export default {
    actions: {
        async loadGroups(ctx: any, filters: {limit: number, component: string}) {
            try {
                const resp = await API.getGroupList({
                    limit: filters.limit,
                    component: filters.component,
                });

                ctx.commit('saveGroups', {active: resp.data.active, fixed: resp.data.fixed});
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        },
    },

    mutations: {
        saveGroups(state: any, groups: { active: Array<openapi.GroupPreview>, fixed: Array<openapi.GroupPreview> }) {
            state.active = groups.active
            state.fixed = groups.fixed
        }
    },

    state: {
        active: Array<openapi.GroupPreview>(),
        fixed: Array<openapi.GroupPreview>(),
    },

    getters: {
        activeGroups(state: any) {
            return state.active
        },

        fixedGroups(state: any) {
            return state.fixed
        },
    },
};