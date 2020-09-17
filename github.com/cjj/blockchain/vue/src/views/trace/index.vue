<template>
    <div class="app-container">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
            <el-form-item label="TraceId" prop="traceId">
                <el-input v-model="queryParams.traceId" placeholder="Please enter TraceId" clearable size="small" @keyup.enter.native="handleQuery" />
            </el-form-item>
            <!-- <el-form-item label="UserAccountId" prop="userAccountId">
                <el-input v-model="queryParams.userAccountId" placeholder="请输入UserAccountId" clearable size="small" @keyup.enter.native="handleQuery" />
            </el-form-item> -->
            
            <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">Search</el-button>
                <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">Empty</el-button>
            </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
            <el-col :span="1.5">
                <el-button
                        type="primary"
                        icon="el-icon-plus"
                        size="mini"
                        @click="handleAdd"
                >Create</el-button>
            </el-col>
            <el-col :span="1.5">
                <el-button
                        icon="el-icon-refresh"
                        size="mini"
                        @click="handleRefresh"
                >Refresh</el-button>
            </el-col>
            <!-- <el-col :span="1.5">
                <el-button
                        type="danger"
                        icon="el-icon-delete"
                        size="mini"
                        :disabled="multiple"
                        @click="handleDelete"
                >删除</el-button>
            </el-col> -->
        </el-row>

        <el-table v-loading="loading" :data="traceList" @selection-change="handleSelectionChange">
            <!-- <el-table-column label="UserAccountId" align="center" prop="userAccountId" :show-overflow-tooltip="true" /> -->
            <el-table-column type="index" width="55" align="center" /><el-table-column label="TraceId" align="center" prop="traceId" :show-overflow-tooltip="true" /><el-table-column label="ProductId" align="center" prop="productId" :show-overflow-tooltip="true" /><el-table-column label="ProductName" align="center" prop="productName" :show-overflow-tooltip="true" /><el-table-column label="ProductNumber" align="center" prop="productNumber" :show-overflow-tooltip="true" /><el-table-column label="ProductPrice" align="center" prop="productPrice" :show-overflow-tooltip="true" /><el-table-column label="ProductTime" align="center" prop="productTime" :show-overflow-tooltip="true" /><el-table-column label="RawIds" align="center" prop="rawIds" :show-overflow-tooltip="true" /><el-table-column label="TraceStatus" align="center" prop="traceStatus" :show-overflow-tooltip="true" />
            <!-- <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                    <el-button
                            size="mini"
                            type="text"
                            icon="el-icon-edit"
                            @click="handleUpdate(scope.row)"
                    >修改</el-button>
                    <el-button
                            size="mini"
                            type="text"
                            icon="el-icon-delete"
                            @click="handleDelete(scope.row)"
                    >删除</el-button>
                </template>
            </el-table-column> -->
        </el-table>

        <!-- <pagination
                v-show="total>0"
                :total="total"
                :page.sync="queryParams.pageIndex"
                :limit.sync="queryParams.pageSize"
                @pagination="getList"
        /> -->

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
            <el-form ref="form" :model="form" :rules="rules" label-width="130px">
                
                <el-form-item label="TraceId" prop="traceId"  >
                <el-input v-model="form.traceId" placeholder="TraceId"  />
                </el-form-item>
                <!-- <el-form-item label="UserAccountId" prop="userAccountId"  >
                <el-input v-model="form.userAccountId" placeholder="UserAccountId"  />
                </el-form-item> -->
                <el-form-item label="ProductId" prop="productId"  >
                <el-input v-model="form.productId" placeholder="ProductId"  />
                </el-form-item>
                <el-form-item label="ProductName" prop="productName"  >
                <el-input v-model="form.productName" placeholder="ProductName"  />
                </el-form-item>
                <el-form-item label="ProductNumber" prop="productNumber"  >
                <el-input v-model="form.productNumber" placeholder="ProductNumber"  />
                </el-form-item>
                <el-form-item label="ProductPrice" prop="productPrice"  >
                <el-input v-model="form.productPrice" placeholder="ProductPrice"  />
                </el-form-item>
                <!-- <el-form-item label="ProductTime" prop="productTime"  >
                <el-input v-model="form.productTime" placeholder="ProductTime"  />
                </el-form-item> -->
                <el-form-item label="RawIds" prop="rawIds"  >
                <el-input v-model="form.rawIds" placeholder="RawIds"  />
                </el-form-item>
                <!-- <el-form-item label="TraceStatus" prop="traceStatus"  >
                <el-input v-model="form.traceStatus" placeholder="TraceStatus"  />
                </el-form-item> -->
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="submitForm">Confirm</el-button>
                <el-button @click="cancel">Cancel</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex'
    import { listTrace, getTrace, delTrace, addTrace, updateTrace } from '@/api/trace'

    export default {
        name: 'Config',
        data() {
            return {
                // 遮罩层
                loading: true,
                // 选中数组
                ids: [],
                // 非单个禁用
                single: true,
                // 非多个禁用
                multiple: true,
                // 总条数
                total: 0,
                // 弹出层标题
                title: '',
                // 是否显示弹出层
                open: false,
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                traceList:[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    traceId:undefined,
                    userAccountId:undefined,
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {traceId:
                    [
                        {required: true, message: 'TraceId不能为空', trigger: 'blur'}
                    ],
                    userAccountId:
                    [
                        {required: true, message: 'UserAccountId不能为空', trigger: 'blur'}
                    ],
                    }
            }
        },
        computed: {
            ...mapGetters([
            'accountId',
            'roles',
            'userName',
            'balance'
        ])
        },
        created() {
            this.getList()
            
        },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listTrace({"userAccountId":this.accountId}).then(response => {
                        this.traceList = response
                        console.log(response)
                        this.loading = false
                    }
                )
            },
            // 取消按钮
            cancel() {
                this.open = false
                this.reset()
            },
            // 表单重置
            reset() {
                this.queryParams.traceId="";
            },
            

            /** 搜索按钮操作 */
            handleQuery() {
                this.loading = true
                getTrace({"userAccountId":this.accountId,"traceId":this.queryParams.traceId}).then(response => {
                        this.traceList = response
                        console.log(response)
                        this.loading = false
                    }
                )
            },
            /** 刷新列表按钮操作 */
            handleRefresh() {
                this.loading = true
                listTrace({"userAccountId":this.accountId}).then(response => {
                        this.traceList = response
                        console.log(response)
                        this.loading = false
                    }
                )
            },

            /** 重置按钮操作 */
            resetQuery() {
                this.queryParams.traceId="";
            },
            /** 新增按钮操作 */
            handleAdd() {
                //this.reset()
                this.open = true
                this.title = 'Create Trace'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.traceId)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 修改按钮操作 */
            handleUpdate(row) {
                this.reset()
                const traceId = row.traceId || this.ids
                getTrace(traceId).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改traceId'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function() {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                    Date.prototype.Format = function (fmt) { // author: meizz
                        var o = {
                            "M+": this.getMonth() + 1, // 月份
                            "d+": this.getDate(), // 日
                            "h+": this.getHours(), // 小时
                            "m+": this.getMinutes(), // 分
                            "s+": this.getSeconds(), // 秒
                            "q+": Math.floor((this.getMonth() + 3) / 3), // 季度
                            "S": this.getMilliseconds() // 毫秒
                        };
                        if (/(y+)/.test(fmt))
                            fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
                        for (var k in o)
                        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
                            return fmt;
                        }
                        var currentDate = new Date().Format("yyyy-MM-dd hh:mm:ss"); 

                        addTrace({"traceId":this.form.traceId,"userAccountId":this.accountId,"productId":this.form.productId,"productName":this.form.productName,"productNumber":parseInt(this.form.productNumber),"productPrice":parseFloat(this.form.productPrice),"productTime":currentDate,"rawIds":this.form.rawIds}).then(response => {
                            alert('Create Trace Success')
                            this.open = false
                            this.getList()
                        })
                        
                    }
                })
            },
            /** 删除按钮操作 */
            handleDelete(row) {
                const Ids = row.traceId || this.ids
                this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(function() {
                    return delTrace(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function() {})
            }
        }
    }
</script>
