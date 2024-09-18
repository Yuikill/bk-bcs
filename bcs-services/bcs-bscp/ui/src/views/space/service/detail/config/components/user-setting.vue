<template>
  <div class="user-settings-wrap">
    <div class="user-content">
      <bk-form-item :label="$t('文件权限')" property="privilege" required>
        <div class="perm-input">
          <bk-popover
            ext-cls="privilege-tips-wrap"
            theme="light"
            trigger="manual"
            placement="top"
            :is-show="showPrivilegeErrorTips">
            <bk-input
              v-model="privilegeInputVal"
              type="number"
              :placeholder="$t('请输入三位权限数字')"
              @blur="handlePrivilegeInputBlur" />
            <template #content>
              <div>{{ t('只能输入三位 0~7 数字') }}</div>
              <div class="privilege-tips-btn-area">
                <bk-button text theme="primary" @click="showPrivilegeErrorTips = false">{{ $t('我知道了') }}</bk-button>
              </div>
            </template>
          </bk-popover>
          <bk-popover ext-cls="privilege-select-popover" theme="light" trigger="click" placement="bottom">
            <div :class="['perm-panel-trigger']">
              <i class="bk-bscp-icon icon-configuration-line"></i>
            </div>
            <template #content>
              <div class="privilege-select-panel">
                <div v-for="(item, index) in PRIVILEGE_GROUPS" class="group-item" :key="index" :label="item">
                  <div class="header">{{ item }}</div>
                  <div class="checkbox-area">
                    <bk-checkbox-group
                      class="group-checkboxs"
                      :model-value="privilegeGroupsValue[index]"
                      @change="handleSelectPrivilege(index, $event)">
                      <bk-checkbox size="small" :label="4" :disabled="index === 0">
                        {{ $t('读') }}
                      </bk-checkbox>
                      <bk-checkbox size="small" :label="2">{{ $t('写') }}</bk-checkbox>
                      <bk-checkbox size="small" :label="1">{{ $t('执行') }}</bk-checkbox>
                    </bk-checkbox-group>
                  </div>
                </div>
              </div>
            </template>
          </bk-popover>
        </div>
      </bk-form-item>
      <div class="user-settings">
        <bk-form-item :label="$t('用户')" property="user" :required="true">
          <bk-select
            v-model="localVal.user"
            :list="userList"
            class="bk-select"
            :filterable="false"
            allow-create
            @select="change">
            <template #optionRender="{ item }">
              <div class="option-item">
                <span>{{ item.label }}</span>
                <span class="bk-bscp-icon icon-close-line close" @click.stop="handleDeleteUser()" />
              </div>
            </template>
          </bk-select>
        </bk-form-item>
        <bk-form-item :label="'UID'" property="user">
          <bk-input v-model="localVal.UID" :placeholder="$t('请输入')" @input="change"></bk-input>
        </bk-form-item>
        <bk-form-item :label="$t('用户组')" :placeholder="$t('请输入')" property="user_group" :required="true">
          <bk-select
            v-model="localVal.user_group"
            :list="userGroupList"
            class="bk-select"
            :filterable="false"
            allow-create
            @select="change">
            <template #optionRender="{ item }">
              <div class="option-item">
                <span>{{ item.label }}</span>
                <span class="bk-bscp-icon icon-close-line close" @click.stop="handleDeleteUserGroup()" />
              </div>
            </template>
          </bk-select>
        </bk-form-item>
        <bk-form-item :label="'GID'" property="user">
          <bk-input v-model="localVal.GID" :placeholder="$t('请输入')" @input="change"></bk-input>
        </bk-form-item>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { useI18n } from 'vue-i18n';

  const { t } = useI18n();

  const showPrivilegeErrorTips = ref(false);
  const privilegeInputVal = ref('');

  const PRIVILEGE_GROUPS = [t('属主（own）'), t('属组（group）'), t('其他人（other）')];
  const PRIVILEGE_VALUE_MAP = {
    0: [],
    1: [1],
    2: [2],
    3: [1, 2],
    4: [4],
    5: [1, 4],
    6: [2, 4],
    7: [1, 2, 4],
  };

  const localVal = ref({
    privilege: '',
    user: '',
    user_group: '',
    UID: 0,
    GID: 0,
  });

  const userList = ref([
    { label: 'root', value: 'root' },
    { label: 'admin', value: 'admin' },
  ]);

  const userGroupList = ref([
    { label: 'root', value: 'root' },
    { label: 'admin', value: 'admin' },
  ]);

  // 权限输入框失焦后，校验输入是否合法，如不合法回退到上次输入
  const handlePrivilegeInputBlur = () => {
    const val = String(privilegeInputVal.value);
    if (/^[0-7]{3}$/.test(val)) {
      localVal.value.privilege = val;
      showPrivilegeErrorTips.value = false;
      change();
    } else {
      privilegeInputVal.value = String(localVal.value.privilege);
      showPrivilegeErrorTips.value = true;
    }
  };

  // 将权限数字拆分成三个分组配置
  const privilegeGroupsValue = computed(() => {
    const data: { [index: string]: number[] } = { 0: [], 1: [], 2: [] };
    if (typeof localVal.value.privilege === 'string' && localVal.value.privilege.length > 0) {
      const valArr = localVal.value.privilege.split('').map((i) => parseInt(i, 10));
      valArr.forEach((item, index) => {
        data[index as keyof typeof data] = PRIVILEGE_VALUE_MAP[item as keyof typeof PRIVILEGE_VALUE_MAP];
      });
    }
    return data;
  });

  // 选择文件权限
  const handleSelectPrivilege = (index: number, val: number[]) => {
    const groupsValue = { ...privilegeGroupsValue.value };
    groupsValue[index] = val;
    const digits = [];
    for (let i = 0; i < 3; i++) {
      let sum = 0;
      if (groupsValue[i].length > 0) {
        sum = groupsValue[i].reduce((acc, crt) => acc + crt, 0);
      }
      digits.push(sum);
    }
    const newVal = digits.join('');
    privilegeInputVal.value = newVal;
    localVal.value.privilege = newVal;
    showPrivilegeErrorTips.value = false;
    change();
  };

  const change = () => {};

  const handleDeleteUser = () => {};

  const handleDeleteUserGroup = () => {};
</script>

<style scoped lang="scss">
  .user-settings-wrap {
    padding: 12px 16px 0px 16px;
    background: #f5f7fa;
    border-radius: 2px;
    .user-content {
      :deep(.bk-form-item) {
        margin-bottom: 16px;
        .bk-form-error {
          position: inherit;
        }
      }
    }
    .user-settings {
      display: flex;
      align-items: center;
      justify-content: space-between;
      :deep(.bk-input) {
        width: 114px;
      }
    }
  }
  .perm-input {
    display: flex;
    align-items: center;
    width: 252px;
    :deep(.bk-input) {
      width: 240px;
      border-right: none;
      border-top-right-radius: 0;
      border-bottom-right-radius: 0;
      .bk-input--number-control {
        display: none;
      }
    }
    .perm-panel-trigger {
      width: 32px;
      height: 32px;
      text-align: center;
      background: #fafcfe;
      color: #3a84ff;
      border: 1px solid #3a84ff;
      cursor: pointer;
      &.disabled {
        color: #dcdee5;
        border-color: #dcdee5;
        cursor: not-allowed;
      }
    }
  }

  .privilege-tips-btn-area {
    margin-top: 8px;
    text-align: right;
  }
  .privilege-select-panel {
    display: flex;
    align-items: top;
    border: 1px solid #dcdee5;
    .group-item {
      .header {
        padding: 0 16px;
        height: 42px;
        line-height: 42px;
        color: #313238;
        font-size: 12px;
        background: #fafbfd;
        border-bottom: 1px solid #dcdee5;
      }
      &:not(:last-of-type) {
        .header,
        .checkbox-area {
          border-right: 1px solid #dcdee5;
        }
      }
    }
    .checkbox-area {
      padding: 10px 16px 12px;
      background: #ffffff;
      &:not(:last-child) {
        border-right: 1px solid #dcdee5;
      }
    }
    .group-checkboxs {
      font-size: 12px;
      .bk-checkbox ~ .bk-checkbox {
        margin-left: 16px;
      }
      :deep(.bk-checkbox-label) {
        font-size: 12px;
      }
    }
  }

  .option-item {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .bk-bscp-icon:hover {
      color: #3a84ff;
    }
  }
</style>

<style></style>
