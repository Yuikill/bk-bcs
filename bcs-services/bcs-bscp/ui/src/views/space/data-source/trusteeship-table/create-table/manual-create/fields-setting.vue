<template>
  <Card :title="$t('字段设置')">
    <template #suffix>
      <div class="add-fields" @click="handleAddFields">
        <Plus class="add-icon" />
        <span class="text">{{ $t('添加字段') }}</span>
      </div>
    </template>
    <bk-table :data="tableData" class="fields-setting-table" :border="['row', 'col', 'outer']" row-hover="auto">
      <bk-table-column :label="$t('字段名')">
        <template #default="{ row }">
          <bk-input v-model="row.fieldsName"></bk-input>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('显示名')" prop="key">
        <template #default="{ row }">
          <bk-input v-model="row.showName"></bk-input>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('数据类型')">
        <template #default="{ row }">
          <bk-select v-if="row" class="type-select" auto-focus :filterable="false" @select="row.type = $event">
            <bk-option v-for="item in dataType" :id="item.value" :key="item.value" :name="item.label" />
          </bk-select>
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('默认值/枚举值')">
        <template #default="{ row }">
          <bk-select v-if="row" class="bk-select" auto-focus filterable @select="row.type = $event">
            <bk-option v-for="item in dataType" :id="item.value" :key="item.value" :name="item.label" />
          </bk-select>
        </template>
      </bk-table-column>
      <!-- <bk-table-column :label="$t('主键')" prop="type"></bk-table-column>
      <bk-table-column :label="$t('非空')" prop="type"></bk-table-column>
      <bk-table-column :label="$t('唯一')" prop="type"></bk-table-column>
      <bk-table-column :label="$t('自增')" prop="type"></bk-table-column>
      <bk-table-column :label="$t('只读')" prop="type"></bk-table-column> -->
      <!-- <bk-table-column></bk-table-column> -->
    </bk-table>
  </Card>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Plus } from 'bkui-vue/lib/icon';
  import Card from '../../../component/card.vue';

  const tableData = ref<any[]>([]);

  const dataType = [
    {
      vakue: 'string',
      label: 'String',
    },
    {
      value: 'number',
      label: 'Number',
    },
    {
      value: 'enum',
      label: 'ENUM',
    },
  ];

  const handleAddFields = () => {
    tableData.value.push({
      fieldsName: '',
      showName: '',
      type: '',
      required: false,
      default: '',
    });
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
  .type-select {
    :deep(.bk-input) {
      height: 41px;
      border: none;
      .bk-input--text {
        padding-left: 16px;
      }
    }
  }
  .fields-setting-table {
    :deep(.bk-table-body) {
      .cell {
        padding: 0;
      }
    }
  }
  .bk-input {
    border: none;
    height: 41px;
    &.is-focused {
      border: 1px solid #3a84ff;
    }
  }
</style>
