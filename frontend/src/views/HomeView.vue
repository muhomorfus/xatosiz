<template>
  <main class="container">
      <br>
      <div class="row">
          <div class="col-lg-3">
              <div class="card mb-3">
                  <div class="card-body">
                      <h4 class="card-title mb-4">Filters</h4>
                      <form class="card-body p-0" @submit.prevent="reload">
                          <div class="form-floating mb-3">
                              <input type="text" class="form-control" id="component" placeholder="service" v-model="component">
                              <label for="component">Component</label>
                          </div>
                          <div class="form-floating mb-3">
                              <input type="number" class="form-control" id="limit" placeholder="10" v-model="limit">
                              <label for="limit">Limit</label>
                          </div>

                          <div class="d-grid gap-1">
                              <button type="submit" class="btn btn-primary">Apply</button>
                          </div>
                      </form>
                  </div>
              </div>
          </div>
          <div class="col-lg-9">
              <h2 class="mb-1">Unfixed events</h2>
              <EventsBar :events="this.allEvents"/>

              <h2 class="mb-3 mt-3">Active groups</h2>
              <GroupPreview v-for="g in activeGroups" :key="g.uuid" :group="g"/>

              <h2 class="mb-3">Fixed groups</h2>
              <GroupPreview v-for="g in fixedGroups" :key="g.uuid" :group="g"/>
          </div>
      </div>
  </main>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {mapActions, mapGetters} from "vuex";
import Group from "@/components/Group.vue";
import EventsBar from "@/components/EventsBar.vue";
import GroupPreview from "@/components/GroupPreview.vue";

export default defineComponent({
    components: {GroupPreview, EventsBar},
    data() {
        return {
            component: '',
            limit: 10,
        }
    },
    methods: {
        ...mapActions(['loadGroups', 'loadEvents']),
        reload() {
            this.loadGroups({limit: this.limit, component: this.component})
        }
    },
    computed: {
        ...mapGetters(['activeGroups', 'fixedGroups', 'allEvents'])
    },
    mounted() {
        this.reload()
        this.loadEvents()
    },
});
</script>