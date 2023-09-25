<template>
  <div class="container w-50">
      <br>

      <div class="card mb-3" v-for="a in alerts" :key="a.uuid">
          <div class="card-header d-flex justify-content-between px-2">
              <Event :event="a.event" style="margin-top: 0rem !important"/>
              <button class="btn btn-sm btn-primary" @click="fixAlert(a.uuid)">Clear</button>
          </div>

          <div class="card-body">
              <h4>{{ a.message }}</h4>

              <small class="text-muted">{{ new Date(a.time).toLocaleString() }}</small>
          </div>
      </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {mapActions, mapGetters} from "vuex";
import Trace from "@/components/Trace.vue";
import EventsBar from "@/components/EventsBar.vue";
import Event from "@/components/Events.vue";

export default defineComponent({
    components: {Event, EventsBar},
    methods: {
        ...mapActions(['loadAlerts', 'fixAlert']),
    },
    computed: {
        ...mapGetters(['alerts']),
    },
    mounted() {
        this.loadAlerts();
    },
});
</script>