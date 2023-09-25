<template>
    <router-link class="simple" :to="`/groups/${group.uuid}`">
        <div class="card mb-3">
            <div class="card-header d-flex justify-content-between">
                <div>
                    {{ group.title }}
                </div>
                <div>
                    {{ group.duration }}
                </div>
            </div>

            <div class="card-body">
                <div class="row">
                    <ComponentsBar style="margin-top: -0.5em" :components="group.components"/>
                    <EventsBar style="margin-top: -0.5em" :events="group.events"/>
                </div>
            </div>
        </div>
    </router-link>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import Trace from "@/components/Trace.vue";
import EventsBar from "@/components/EventsBar.vue";
import ComponentsBar from "@/components/ComponentsBar.vue";

export default defineComponent({
    components: {ComponentsBar, EventsBar},
    props: ['group'],
});
</script>

<style scoped>
.simple {
    text-decoration: none;
    color: black;
}
</style>