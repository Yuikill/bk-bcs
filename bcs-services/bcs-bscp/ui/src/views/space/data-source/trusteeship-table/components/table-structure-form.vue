<template>
  <div class="table-structure-form">
    <Card :title="$t('字段设置')">
      <template v-if="isManualCreate" #suffix>
        <div class="add-fields" @click="handleAddFields">
          <Plus class="add-icon" />
          <span class="text">{{ $t('添加字段') }}</span>
        </div>
      </template>
      <FieldsTable v-if="isManualCreate" :list="filedsList" />
      <UploadFieldsTable v-else-if="filedsList.length" :list="filedsList"></UploadFieldsTable>
      <bk-exception
        v-else
        class="exception-wrap-item"
        :description="$t('请先上传文件')"
        :title="$t('暂无数据')"
        type="empty" />
    </Card>
    <bk-form form-type="vertical" :model="formData">
      <Card :title="$t('基本信息')">
        <div class="basic-info-form">
          <bk-form-item :label="$t('表格名称')" required>
            <bk-input v-model="formData.name" :disabled="isEdit"></bk-input>
          </bk-form-item>
          <bk-form-item :label="$t('表格描述')">
            <bk-input v-model="formData.memo"></bk-input>
          </bk-form-item>
        </div>
      </Card>
      <Card :title="$t('可见范围')">
        <bk-form-item :label="$t('选择服务')" required>
          <bk-select
            v-model="formData.bound_apps"
            :loading="serviceLoading"
            style="width: 464px"
            multiple
            filterable
            :placeholder="$t('请选择服务')">
            <bk-option :label="$t('全部服务')" :value="0"></bk-option>
            <bk-option v-for="service in serviceList" :key="service.id" :label="service.spec.name" :value="service.id">
            </bk-option>
          </bk-select>
        </bk-form-item>
      </Card>
    </bk-form>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { getAppList } from '../../../../../api/index';
  import { IAppItem } from '../../../../../../types/app';
  import { IFiledsItem } from '../../../../../../types/kv-table';
  import { Plus } from 'bkui-vue/lib/icon';
  import Card from '../../component/card.vue';
  import FieldsTable from './fields-table.vue';
  import UploadFieldsTable from './upload-fields-table.vue';

  const props = defineProps<{
    bkBizId: string;
    isManualCreate: boolean;
    isEdit: boolean;
  }>();

  const formData = ref({
    name: '',
    memo: '',
    bound_apps: [],
  });
  const serviceLoading = ref(false);
  const serviceList = ref<IAppItem[]>([]);

  const filedsList = ref<IFiledsItem[]>([
    {
      fieldsName: '123',
      showName: '123',
      type: '',
      required: false,
      default: '',
      primaryKey: true,
      nonEmpty: false,
      only: false,
      autoIncrement: false,
      readonly: false,
      isShowSettingEnumPopover: false,
    },
  ]);

  onMounted(() => {
    getServiceList();
  });

  const handleAddFields = () => {
    const newField = {
      fieldsName: '',
      showName: '',
      type: '',
      required: false,
      default: '',
      primaryKey: filedsList.value.length === 0,
      nonEmpty: false,
      only: false,
      autoIncrement: false,
      readonly: false,
      isShowSettingEnumPopover: false,
    };

    filedsList.value.push(newField);
  };

  const getServiceList = async () => {
    serviceLoading.value = true;
    try {
      const bizId = props.bkBizId;
      const query = {
        start: 0,
        limit: 1000, // @todo 确认拉全量列表参数
      };
      const resp = await getAppList(bizId, query);
      serviceList.value = resp.details;
    } catch (e) {
      console.error(e);
    } finally {
      serviceLoading.value = false;
    }
  };
</script>

<style scoped lang="scss">
  .add-fields {
    display: flex;
    align-items: center;
    height: 16px;
    cursor: pointer;
    .add-icon {
      border-radius: 50%;
      background-color: #3a84ff;
      color: #fff;
      margin-right: 5px;
    }
    .text {
      color: #3a84ff;
      font-size: 12px;
    }
  }

  .table-structure-form {
    .card:not(:last-child) {
      margin-bottom: 16px;
    }
  }
  .basic-info-form {
    display: flex;
    gap: 24px;
    .bk-form-item {
      flex: 1;
    }
  }

  .exception-wrap-item {
    :deep(.bk-exception-img) {
      width: 280px;
      height: 140px;
    }
    :deep(.bk-exception-title) {
      margin-top: 8px;
      font-size: 14px;
      color: #63656e;
      line-height: 22px;
    }
    :deep(.bk-exception-description) {
      margin-top: 8px;
      font-size: 12px;
      color: #979ba5;
      line-height: 20px;
    }
  }
</style>
