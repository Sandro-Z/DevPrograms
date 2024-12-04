<template>
	<div class = "main_conect"> 	
		<div class = "show_main">
			<router-view></router-view>
		</div>
	</div>
</template>

<script>
  export default{
    name: 'TreeMenu',
	data(){
		return{
			funcs:[
				{text:'个人工作区',Isselected:false,id:"1",
					details:[{id:"1",text:'个人设置',link:'set_of_member'}]},
				{text:'设备与端口',Isselected:false,id:"2",
					details:[{id:"2",text:'端口查询',link:'search'},{id:"3",text:'设备监控',link:'state'}]},
				{text:'网络状态图',Isselected:false,id:"3",
					details:[{text:'网络流量图',link:''}]}
			]
		}
	},
    methods: {
		changeSelected(index){
			this.funcs[index].Isselected = !this.funcs[index].Isselected;
		}
	}
  }
</script>

<style>
  
</style>
