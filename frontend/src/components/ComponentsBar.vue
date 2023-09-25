<template>
    <div class="col-auto">
        <button v-for="c in components" :key="c.name"
           :class="`btn btn-sm btn-light mt-2 me-2`">
            {{ c.name }} ({{ c.quantity }})
        </button>
    </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
    name: 'ComponentsBar',
    props: ['components']
});
</script>