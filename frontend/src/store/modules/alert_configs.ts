import Vuex from 'vuex';

import * as openapi from '@/openapi';

import API from '@/api';

export default {
    actions: {
        async loadAlertConfigs(ctx: any) {
            try {
                const resp = await API.getAlertConfigList();

                ctx.commit('saveAlertConfigs', resp.data.items);
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        },

        async newAlertConfig(ctx: any, cfg: {
            message_expression: string,
            min_priority: string,
            duration: string,
            min_rate: number,
            comment: string,
        }) {
            try {
                const resp = await API.createAlertConfig({
                    message_expression: cfg.message_expression,
                    min_priority: cfg.min_priority,
                    duration: cfg.duration,
                    min_rate: cfg.min_rate,
                    comment: cfg.comment,
                });

                ctx.commit('addAlertConfig', resp.data);
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        },

        async deleteAlertConfig(ctx: any, id: string) {
            try {
                await API.deleteAlertConfig(id);

                ctx.commit('removeAlertConfig', id);
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        }
    },

    mutations: {
        saveAlertConfigs(state: any, configs: Array<openapi.AlertConfig>) {
            state.configs = configs;
        },

        addAlertConfig(state: any, config: openapi.AlertConfig) {
            state.configs.push(config)
        },

        removeAlertConfig(state: any, id: string) {
            state.configs = state.configs.filter((value: openapi.AlertConfig) => value.uuid != id)
        }
    },

    state: {
        configs: Array<openapi.AlertConfig>(),
    },

    getters: {
        alertConfigs(state: any) {
            return state.configs
        },
    },
};