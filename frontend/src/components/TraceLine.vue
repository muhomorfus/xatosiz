<template>
    <div class="progress" style="height: 8px">
        <div class="progress-bar bg-white" role="progressbar" style="width: 100%" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100"></div>
    </div>

    <div class="progress" style="height: 8px">
        <div class="progress-bar bg-transparent" role="progressbar" :style="`width: ${start}%`" :aria-valuenow="start" :aria-valuemin="0" aria-valuemax="100"></div>
        <div class="progress-bar bg-success" role="progressbar" :style="`width: ${end - start}%`" :aria-valuenow="end - start" :aria-valuemin="0" aria-valuemax="100"></div>
    </div>

    <div class="progress" style="height: 8px">
        <div class="progress-bar bg-white" role="progressbar" style="width: 100%" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100"></div>
    </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
    props: ['start', 'end'],
});
</script>

<style scoped>

</style>