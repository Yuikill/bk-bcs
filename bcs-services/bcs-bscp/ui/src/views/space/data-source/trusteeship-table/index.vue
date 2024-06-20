<template>
  <div class="operate-area">
    <div class="operate-btns">
      <bk-button theme="primary" @click="handleOpenCreate">{{ $t('新建数据源') }}</bk-button>
      <div class="tab">
        <div
          v-for="item in panels"
          :key="item.value"
          :class="['tab-item', { active: active === item.value }]"
          v-bk-tooltips="{ content: $t('敬请期待'), disabled: item.value !== 'tencent-doct' }"
          @click="handleTabClick(item.value)">
          {{ item.label }}
        </div>
      </div>
    </div>
    <SearchInput
      v-model="searchStr"
      class="config-search-input"
      :width="280"
      :placeholder="$t('数据源名称/数据源别名/最近更新人')" />
  </div>
  <Table />
  <CreateTable v-if="showCreateTable" @close="showCreateTable = false"/>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import SearchInput from '../../../../components/search-input.vue';
  import Table from './table.vue';
  import CreateTable from './create-table/index.vue';

  const searchStr = ref('');
  const active = ref('all');
  const panels = [
    {
      value: 'all',
      label: '全部',
    },
    {
      value: 'mysql',
      label: 'MySQL',
    },
    {
      value: 'tencent-doct',
      label: '腾讯文档',
    },
  ];
  const showCreateTable = ref(false);

  const handleTabClick = (value: string) => {
    if (value === 'tencent-doct') return;
    active.value = value;
  };

  const handleOpenCreate = () => {
    showCreateTable.value = true;
  };
</script>

<style scoped lang="scss">
  .operate-area {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 16px;
    .operate-btns {
      display: flex;
      align-items: center;
      .tab {
        display: flex;
        height: 32px;
        background: #f0f1f5;
        border-radius: 2px;
        margin-left: 16px;
        padding: 4px;
        .tab-item {
          padding: 1px 12px;
          display: flex;
          align-items: center;
          justify-content: center;
          height: 24px;
          color: #63656e;
          cursor: pointer;
          &.active {
            background: #fff;
            color: #3a84ff;
          }
        }
      }
    }
  }
</style>
