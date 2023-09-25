<template>
  <div class="container w-75">
      <br>
      <div class="row">
          <div class="col-lg-3">
              <div class="card">
                  <ul class="list-group list-group-flush">
                      <li class="list-group-item">
                          <router-link style="text-decoration: none" to="/settings/alerts">Alerts</router-link>
                      </li>
                  </ul>
              </div>
          </div>

          <div class="col-lg-9">
              <div class="card mb-3">
                  <div class="card-header">
                      Add alert
                  </div>

                  <div class="card-body">
                      <form class="card-body p-0" @submit.prevent="addConfig">
                          <div class="form-floating mb-3">
                              <input type="text" class="form-control" id="regexp" placeholder="*" v-model="regexp">
                              <label for="regexp">Regular expression</label>
                          </div>

                          <div class="form-floating mb-3">
                              <select name="priority" id="priority" class="form-control" v-model="priority">
                                  <option>info</option>
                                  <option>warning</option>
                                  <option>error</option>
                                  <option>fatal</option>
                              </select>

                              <label for="priority">Priority</label>
                          </div>

                          <div class="form-floating mb-3">
                              <input type="text" class="form-control" id="duration" placeholder="10m" v-model="duration">
                              <label for="duration">Duration</label>
                          </div>

                          <div class="form-floating mb-3">
                              <input type="number" class="form-control" id="rate" placeholder="10" v-model="rate">
                              <label for="rate">Rate</label>
                          </div>

                          <div class="form-floating mb-3">
                              <input type="text" class="form-control" id="comment" placeholder="10m" v-model="comment">
                              <label for="comment">Comment</label>
                          </div>

                          <div class="d-grid gap-1">
                              <button type="submit" class="btn btn-primary">Apply</button>
                          </div>
                      </form>
                  </div>
              </div>

              <div class="card mb-3" v-for="c in alertConfigs" :key="c.uuid">
                  <div class="card-header d-flex justify-content-between px-2">
                      <div></div>
                      <button class="btn btn-sm btn-danger" @click="deleteAlertConfig(c.uuid)">Delete</button>
                  </div>

                  <ul class="list-group list-group-flush">
                      <li class="list-group-item">
                          <strong>Regexp: </strong> {{ c.message_expression }}
                      </li>

                      <li class="list-group-item">
                          <strong>Min priority: </strong> {{ c.min_priority }}
                      </li>

                      <li class="list-group-item">
                          <strong>Duration: </strong> {{ c.duration }}
                      </li>

                      <li class="list-group-item">
                          <strong>Min rate: </strong> {{ c.min_rate }}
                      </li>

                      <li class="list-group-item">
                          <strong>Notification text: </strong> {{ c.comment }}
                      </li>
                  </ul>
              </div>
          </div>
      </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {mapActions, mapGetters} from "vuex";
import Trace from "@/components/Trace.vue";

export default defineComponent({
    data() {
        return {
            regexp: '',
            priority: '',
            duration: '',
            rate: 0,
            comment: '',
        }
    },
    methods: {
        ...mapActions(['loadAlertConfigs', 'newAlertConfig', 'deleteAlertConfig']),
        addConfig() {
            this.newAlertConfig({
                message_expression: this.regexp,
                min_priority: this.priority,
                duration: this.duration,
                min_rate: this.rate,
                comment: this.comment,
            });
        }
    },
    computed: {
        ...mapGetters(['alertConfigs']),
    },
    mounted() {
        this.loadAlertConfigs();
    },
});
</script>