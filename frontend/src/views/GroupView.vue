<template>
  <div class="container">
      <br>
      <Group :group="getGroup"/>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {mapActions, mapGetters} from "vuex";
import Group from "@/components/Group.vue";

export default defineComponent({
    components: {Group},
    methods: {
        ...mapActions(['loadGroup'])
    },
    computed: {
        ...mapGetters(['getGroup']),
    },
    created() {
        this.loadGroup(this.$route.params.uuid.toString())
    },
});
</script>