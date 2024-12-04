<template>
	<div class="container-fluid bg-primary top">
		<div class = "row">
			<div class = "col-md-4">
			</div>
			<div class = "col-md-4 text-center color_title">
				<h2>西安交通大学网络管理系统</h2>
			</div>
			<div class = "col-md-1 text-center">
				<user_set></user_set>
			</div>
			<div class = "col-md-1 text-center">
				<search_set></search_set>
			</div>
		</div>
	</div>
</template>

<script>
import user_set from './Main/user_set.vue'
import search_set from './Main/search_key'

export default{
	name:'top',

	data() {
		return {
				user:{name:'郑新宇'},
			};
	},
	methods:{
		
	},
	components:{
		user_set,
		search_set
	}
} 
</script>

<style>
</style>
