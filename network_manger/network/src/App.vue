<template>
	<div id='app'>
		<div>
			<top></top>
		</div>
		<div>
			<main_context></main_context>
		</div>
	</div>
	
</template>

<script>
import top from './components/top.vue'
import main_context from './components/main.vue'

export default {
  name: 'App',
  components:{
		top,
		main_context
	}
}
</script>

<style>
</style>
