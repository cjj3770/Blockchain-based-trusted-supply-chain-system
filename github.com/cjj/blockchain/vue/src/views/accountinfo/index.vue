<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>Account ID: {{ accountId }}</p>
      <p>Username: {{ userName }}</p>
      <p>Organization ID: {{ serviceId }}</p>

    </el-alert>
    <div v-if="organizationInfoList.length==0" style="text-align: center;">
      <el-alert
        title="cannot find data"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in organizationInfoList" :key="index" :span="6" :offset="1">
        <el-card class="realEstate-card">
          <div slot="header" class="clearfix">
            Organization Information
          </div>

          <div class="item">
            <el-tag>ID: </el-tag>
            <span>{{ val.Id }}</span>
          </div>
          <div class="item">
            <el-tag type="success">Name: </el-tag>
            <span>{{ val.Name }}</span>
          </div>
          <div class="item">
            <el-tag type="success">Address: </el-tag>
            <span>{{ val.Address }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">Funds: </el-tag>
            <span>{{ val.funds }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">Principal: </el-tag>
            <span>{{ val.Principal }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getLogisticCom,getManufacturer,getStore} from '@/api/organization'

export default {
  name: 'organization',
  data() {
    return {
      loading: true,
      organizationInfoList: [],
      accountList: [],
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'roles',
      'userName',
      'serviceId',
      'balance'
    ])
  },
  created() {
    if (this.roles[0] === 'admin'||this.roles[0] === 'courier') {
      getLogisticCom({"Id":this.serviceId}).then(response => {
        if (response !== null) {
          this.organizationInfoList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else if (this.roles[0] === 'store') {
      getStore({"Id":this.serviceId}).then(response => {
        if (response !== null) {
          this.organizationInfoList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }else if (this.roles[0] === 'manuf') {
      getManufacturer({"Id":this.serviceId}).then(response => {
        console.log(response)
        if (response !== null) {
          this.organizationInfoList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },


}

</script>

<style>
  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }
  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .realEstate-card {
    width: 280px;
    height: 340px;
    margin: 18px;
  }
</style>
