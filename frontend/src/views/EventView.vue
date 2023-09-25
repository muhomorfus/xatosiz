<template>
  <div class="container w-50">
      <br>
      <div :class="`card text-white ${this.bg} mb-3`">
          <div class="card-header d-flex justify-content-between">
              <div>
                  {{ this.title }}
              </div>
              <div>
                  <a v-if="!this.event.fixed" class="fix" href="#" @click.prevent="this.fix">Fix this event</a>
              </div>
          </div>

          <div class="card-body">
              <h5 class="card-title">{{ this.event.message }}</h5>
              <div class="card-text code styled-scrollbars">{{ this.payload }}</div>
          </div>
          <div class="card-footer">
              <small>{{ new Date(this.event.time).toLocaleString() }}</small>
          </div>
      </div>
  </div>
</template>

<style>
.code {
    font-family: monospace;
    border: 1px solid var(--bs-border-color-translucent);
    border-radius: 7px;
    padding: 10px;
    white-space: pre;
    overflow: scroll;
}

.fix {
    color: white;
    font-weight: bold;
}
.fix:hover,
.fix:focus,
.fix:active {
    cursor: pointer;
    color: white;
}

.styled-scrollbars {
    --scrollbar-foreground: transparent;
    --scrollbar-background: transparent;
    --scrollbar-size: 0px;
    scrollbar-color: var(--scrollbar-foreground) var(--scrollbar-background);
}
.styled-scrollbars::-webkit-scrollbar {
    width: var(--scrollbar-size);
    height: var(--scrollbar-size);
}
.styled-scrollbars::-webkit-scrollbar-thumb {
    background: var(--scrollbar-foreground);
}
.styled-scrollbars::-webkit-scrollbar-track {
    background: var(--scrollbar-background);
}
</style>

<script lang="ts">
import {defineComponent} from "vue";
import {mapActions, mapGetters} from "vuex";

export default defineComponent({
    methods: {
        ...mapActions(['loadEvent', 'fix'])
    },
    computed: {
        ...mapGetters(['event']),
        title(): string {
            return this.event.priority
        },

        bg(): string {
            switch (this.event.priority) {
                case 'info':
                    return 'bg-primary'
                case 'warning':
                    return 'bg-warning'
                case 'error':
                    return 'bg-danger'
                case 'fatal':
                    return 'bg-dark'
                default:
                    return 'bg-primary'
            }
        },

        payload(): string {
            return JSON.stringify(this.event.payload, null, 2)
        }
    },
    mounted() {
        this.loadEvent(this.$route.params.uuid.toString())
    },
});
</script>