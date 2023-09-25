<template>
    <li class="list-group-item">
        <div class="row trace">
            <div class="col-lg-3 col-md-4 styled-scrollbars" :style="`padding-left: ${12 + trace.level * 20}px; overflow: scroll`">
                <b style="white-space: nowrap;">{{ trace.trace.title }} &nbsp;</b>
                <i class="text-muted" style="white-space: nowrap;">{{ trace.trace.component }}</i>
            </div>
            <div class="col-lg-9 col-md-8">
                <TraceLine :start="percent()(trace.trace.start)"
                           :end="percent()(trace.trace.end)"/>
            </div>
        </div>

        <div v-if="trace.trace.events.length" class="row"
             :style="`padding-left: ${trace.level * 20}px`">
            <EventsBar :events="trace.trace.events"/>
        </div>
    </li>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import TraceLine from "@/components/TraceLine.vue";
import {mapGetters} from "vuex";
import EventsBar from "@/components/EventsBar.vue";

export default defineComponent({
    name: 'Trace',
    components: {EventsBar, TraceLine},
    props: ['trace'],
    methods: {
        ...mapGetters(['percent']),
        buttonClass(priority: string): string {
            switch (priority) {
                case 'info':
                    return 'btn-outline-primary'
                case 'warning':
                    return 'btn-outline-warning'
                case 'error':
                    return 'btn-outline-danger'
                case 'fatal':
                    return 'btn-danger'
                default:
                    return 'btn-outline-primary'
            }
        }
    }
});
</script>

<style>
.styled-scrollbars {
    --scrollbar-foreground: transparent;
    --scrollbar-background: transparent;
    --scrollbar-size: 0px;
    /* плашка-бегунок, фон */
    scrollbar-color: var(--scrollbar-foreground) var(--scrollbar-background);
}
.styled-scrollbars::-webkit-scrollbar {
    width: var(--scrollbar-size); /* в основном для вертикальных полос прокрутки */
    height: var(--scrollbar-size); /* в основном для горизонтальных полос прокрутки */
}
.styled-scrollbars::-webkit-scrollbar-thumb { /* плашка-бегунок */
    background: var(--scrollbar-foreground); /* фон */
}
.styled-scrollbars::-webkit-scrollbar-track { /* фон */
    background: var(--scrollbar-background);
}
</style>