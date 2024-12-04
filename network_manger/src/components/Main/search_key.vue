<template>
	<div>
		<button type="button" class="btn btn-info" 
                data-toggle="modal" data-target="#search_way" data-whatever="@mdo">端口查询</button>

        <div class="modal fade" id="search_way" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <h4 class="modal-title" id="exampleModalLabel">端口查询</h4>
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    </div>
                    <div class="modal-body">
                        <form class="form-horizontal">
                            <div class="form-group dis_flex container-fluid" v-for="key_s in key_ss" :key = "key_s.id">
                                <label for="key_s.text" class="control-label col-sm-4">{{key_s.text}}</label>
                                <input type="text" class="form-control col-sm-8" id="recipient-name">
                            </div>
                        </form>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                        <button type="button" class="btn btn-primary">查找</button>
                    </div>
                </div>
            </div>
        </div>
	</div>
</template>

<script>
	export default{
		name:'search_key',
		data(){
			return{
				key_ss:[
					{text:'用户住址'},
					{text:'用户姓名'},
					{text:'用户IP地址'},
					{text:'交换机IP地址'},
					{text:'交换机端口号'},
				]
			}
		}
	}
</script>

<style>
</style>
