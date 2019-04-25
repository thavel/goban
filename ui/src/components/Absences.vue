<template>
<div>
  <month :year="2019" :month="3" :header="true"/>
  <month :year="2019" :month="3"/>
</div>
</template>

<script>
import Month from '@/components/ui/Month.vue';

export default {

  components: {
    Month
  }
}
</script>

<style>
</style>
