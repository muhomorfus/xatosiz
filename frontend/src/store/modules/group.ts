import Vuex from 'vuex';

import * as openapi from '@/openapi';

import API from '@/api';

function timeToNumber(time: string): number {
    return new Date(time).valueOf();
}

function minTraceTime(trace: openapi.Trace): number {
    let min: number = timeToNumber(trace.start);

    for (let i = 0; i < trace.children.length; i++) {
        const localMin = minTraceTime(trace.children[i]);

        if (localMin < min) {
            min = localMin;
        }
    }

    return min;
}

function maxTraceTime(trace: openapi.Trace): number {
    let max: number = timeToNumber(trace.end);

    for (let i = 0; i < trace.children.length; i++) {
        const localMax = maxTraceTime(trace.children[i]);

        if (localMax > max) {
            max = localMax;
        }
    }

    return max;
}

function toOneArray(traces: Array<openapi.Trace>): Array<{level: number, trace: openapi.Trace}> {
    let result = traces.map((t: openapi.Trace) => {
        return {
            level: 0,
            trace: t,
        }
    });
    let finish = false;

    while (!finish) {
        finish = true;

        const newResult = Array<{level: number, trace: openapi.Trace}>();
        for (let i = 0; i < result.length; i++) {
            const children = result[i].trace.children;

            result[i].trace.children = [];
            newResult.push(result[i]);

            if (children.length) {
                finish = false;
            }

            for (let j = 0; j < children.length; j++) {
                newResult.push({
                    level: result[i].level + 1,
                    trace: children[j],
                });
            }
        }

        result = newResult;
    }

    return result
}

function minTime(group: openapi.Group): number {
    return Math.min(...group.traces.map(t => minTraceTime(t)));
}

function maxTime(group: openapi.Group): number {
    return Math.max(...group.traces.map(t => maxTraceTime(t)));
}

export default {
    actions: {
        async loadGroup(ctx: any, uuid: string) {
            try {
                const resp = await API.getGroup(uuid);

                console.error(resp.data);
                ctx.commit('saveGroup', resp.data);
            } catch (e) {
                console.error(e);
                ctx.dispatch('showError', e.response.data.comment);
            }
        },
    },

    mutations: {
        saveGroup(state: any, g: openapi.Group) {
            state.group = g
            state.minTime = minTime(g)
            state.maxTime = maxTime(g)
        },
    },

    state: {
        group: {} as openapi.Group,
        maxTime: 0,
        minTime: 0,
    },

    getters: {
        getGroup(state: any) {
            return {
                uuid: state.group.uuid,
                traces: toOneArray(state.group.traces),
            }
        },

        percent(state: any) {
            return function(timeStr: string) {
                let time = timeToNumber(timeStr);

                time -= state.minTime;
                time *= 100;
                time /= (state.maxTime - state.minTime);

                return time;
            }
        },
    },
};