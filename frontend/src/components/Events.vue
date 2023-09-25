<template>
    <a
       :class="`btn btn-sm ${buttonClass(event.priority)} mt-2 me-2`"
       :href="`/events/${event.uuid}`">
        {{ event.message }}
    </a>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import TraceLine from "@/components/TraceLine.vue";
import {mapGetters} from "vuex";

export default defineComponent({
    name: 'Event',
    props: ['event'],
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