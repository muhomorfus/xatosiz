<template>
    <div class="col-auto">
        <div class="col-auto">
            <Event v-for="e in events" :key="e.id" :event="e"/>
        </div>
    </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import TraceLine from "@/components/TraceLine.vue";
import {mapGetters} from "vuex";
import Event from "@/components/Events.vue";

export default defineComponent({
    name: 'EventsBar',
    components: {Event},
    props: ['events'],
    methods: {
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