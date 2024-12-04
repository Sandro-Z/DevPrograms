<template>
	<div class = "state">
        <div class = "container-fluid top_button">
            <div class = "row">
                <div class = "col-md-1"></div>
                <button type="button" class="btn btn-success nomarl col-md-4 btn-lg" 
                        @click="change_show(true)">东校区学生宿舍</button>
                <div class = "col-md-2"></div>
                <button type="button" class="btn btn-success nomarl col-md-4 btn-lg" 
                        @click="change_show(false)">西校区学生宿舍</button>
                <div class = "col-md-1"></div>
            </div>
        </div>
        <div v-show="east">
            <table class = "table table-bordered wid_80">
                <tr class = "row">
                    <td v-for = "index in indexs" :key="index.id" :class="index.state?'nomarl':'wrong'"
                        class = "col-md-3 text-center">{{index.name}}</td>
                </tr>
            </table>
        </div>
       <div v-show="!east">
            <table class = "table table-bordered wid_80">
                <tr class = "row">
                    <td v-for = "index in indexs_w" :key="index.id" :class="index.state?'nomarl':'wrong'"
                        class = "col-md-3 text-center">{{index.name}}</td>
                </tr>
            </table>
        </div>
    </div>
</template> 

<script>
	export default{
		name:'search_key',
		data(){
			return{
                east:true,
				key_ss:[
					{text:'用'},
					{text:'用户姓名'},
					{text:'用户IP地址'},
					{text:'交换机IP地址'},
					{text:'交换机端口号'},
                ],
                indexs:[
                    {name:'二区学1号楼',id:'1',state:false},
                    {name:'二区学2号楼',id:'2',state:true},
                    {name:'二区学3号楼',id:'3',state:true},
                    {name:'二区学4号楼',id:'4',state:false},
                    {name:'团委学生科技社团中心',id:'5',state:true},
                    {name:'学东10舍',id:'6',state:true},
                    {name:'学东11舍',id:'7',state:true},
                    {name:'学东12舍',id:'8',state:true},
                    {name:'学东13舍',id:'9',state:true},
                    {name:'学东1、4舍',id:'10',state:true},
                    {name:'学东14舍',id:'11',state:true},
                    {name:'学东15舍',id:'12',state:true},
                    {name:'学东16舍',id:'13',state:true},
                    {name:'学东17舍',id:'14',state:true},
                    {name:'学东18、19、20（1-4）',id:'15',state:true},
                    {name:'学东18、19、20（5-8）',id:'16',state:true},
                    {name:'学东22舍',id:'17',state:true},
                    {name:'学东2、3舍',id:'',state:false},
                    {name:'学东23舍',id:'',state:true},
                    {name:'学东6舍',id:'',state:true},
                    {name:'学东7舍',id:'',state:true},
                    {name:'学东8舍',id:'',state:true},
                    {name:'学东9舍',id:'',state:true},
                    {name:'学南2号楼',id:'',state:true},
                    {name:'学生区PPPOE',id:'',state:true},
                    {name:'学西11舍',id:'',state:true},
                    {name:'学西12舍',id:'',state:true},
                    {name:'学西13舍',id:'',state:true},
                    {name:'学西14舍',id:'',state:true},
                    {name:'学西15舍',id:'',state:true},
                    {name:'学西16舍',id:'',state:true},
                    {name:'学西17舍',id:'',state:true},
                    {name:'学西18、19、20（1-4）舍',id:'',state:true},
                    {name:'学西18、19、20（5-8）舍',id:'',state:true},
                    {name:'学西1舍',id:'',state:true},
                    {name:'学西21-23舍',id:'',state:true},
                    {name:'学西2舍',id:'',state:true},
                    {name:'学西3舍',id:'',state:true},
                    {name:'学西4舍',id:'',state:true},
                    {name:'学西5舍',id:'',state:true},
                    {name:'学西6舍',id:'',state:true},
                    {name:'学西8舍',id:'',state:true},
                    {name:'学西9、10舍',id:'',state:true},
                ],
                indexs_w:[
                    {name:'财经留学楼',id:'1',state:true},
                    {name:'二附院学生2号楼',id:'1',state:true},
                    {name:'二附院学生宿舍',id:'1',state:true},
                    {name:'西区wes4',id:'1',state:true},
                    {name:'西区wesUe',id:'1',state:true},
                    {name:'西区wesUn',id:'1',state:true},
                    {name:'西区wesUs',id:'1',state:true},
                    {name:'西区wmsl',id:'1',state:true},
                    {name:'西区wms10-101112',id:'1',state:true},
                    {name:'西区wms10-123',id:'1',state:true},
                    {name:'西区wms10-456',id:'1',state:true},
                    {name:'西区wms10-789',id:'1',state:true},
                    {name:'西区wms2',id:'1',state:false},
                    {name:'西区wms3',id:'1',state:false},
                    {name:'西区wms5',id:'1',state:true},
                    {name:'西区wms6',id:'1',state:true},
                    {name:'西区wms7',id:'1',state:true},
                    {name:'西区wms8',id:'1',state:true},
                    {name:'西区医学4号楼',id:'1',state:true},
                ]
			}
        },
        methods:{
            change_show(bo){
                this.east = bo;
            }
        }
	}
</script>

<style>
</style>
