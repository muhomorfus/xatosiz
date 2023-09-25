<template>
    <div class="card mb-3" v-if="group.traces.length">
        <ul class="list-group list-group-flush">
            <Trace v-for="t in group.traces" :key="t.uuid"
                   :trace="t"/>
        </ul>
    </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import Trace from "@/components/Trace.vue";

export default defineComponent({
    components: {Trace},
    props: ['group'],
});
</script>

<style scoped>
</style>