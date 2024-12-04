<template>
   <div class = "user_information">
        <button type="button" class="btn btn-info" 
                data-toggle="modal" data-target="#exampleModal" data-whatever="@mdo">个人信息</button>

        <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <h4 class="modal-title" id="exampleModalLabel">个人信息设置</h4>
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    </div>
                    <div class="modal-body">
                        <form class="form-horizontal">
                            <div class="form-group dis_flex container-fluid" v-for="key_s in key_ss" :key = "key_s.id">
                                <label for="key_s.text" class="control-label col-sm-2">{{key_s.text}}</label>
                                <input type="text" class="form-control col-sm-10" id="recipient-name">
                            </div>
                        </form>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                        <button type="button" class="btn btn-primary">提交修改</button>
                    </div>
                </div>
            </div>
        </div>
   </div>
</template>

<script>
export default {
    name:'user_set',
    data(){
        return{
            key_ss:[
					{id:'1',text:'姓名'},
					{id:'2',text:'部门'},
					{id:'3',text:'办公地点'},
					{id:'4',text:'固定电话'},
                    {id:'5',text:'手机'},
                    {id:'6',text:'电子信箱'},
                    {id:'7',text:'账号角色'}
                ]
        }
    },
    methods:{
    }
}
</script>