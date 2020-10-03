// https://umijs.org/config/
import { defineConfig } from 'umi';
import defaultSettings from './defaultSettings';
import proxy from './proxy';
const { REACT_APP_ENV } = process.env;
export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  layout: {
    name: 'Ant Design Pro',
    locale: true,
  },
  locale: {
    // default zh-CN
    default: 'zh-CN',
    antd: true,
    // default true, when it is true, will use `navigator.language` overwrite default
    baseNavigator: true,
  },
  dynamicImport: {
    loading: '@/components/PageLoading/index',
  },
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/docs/routing
  routes: [
    {
      path: '/user',
      layout: false,
      routes: [
        {
          name: 'login',
          path: '/user/login',
          component: './user/login',
        },
      ],
    },
    {
      path: '/helm',
      name: 'helm',
      icon: 'crown',
      routes: [
        {
          name: 'release',
          path: '/helm/release',
          component: './Welcome',
        },
      ],
    },
    {
      name: '错误页',
      icon: 'smile',
      path: '/exception',
      routes: [
        {
          name: '403',
          icon: 'smile',
          path: '/exception/403',
          component: './Exception403',
        },
        {
          name: '404',
          icon: 'smile',
          path: '/exception/404',
          component: './Exception404',
        },
        {
          name: '500',
          icon: 'smile',
          path: '/exception/500',
          component: './Exception500',
        },
      ],
    },
    {
      path: '/welcome',
      name: 'welcome',
      icon: 'smile',
      component: './Welcome',
    },
    {
      path: '/admin',
      name: 'admin',
      icon: 'crown',
      access: 'canAdmin',
      component: './Admin',
      routes: [
        {
          path: '/admin/sub-page',
          name: 'sub-page',
          icon: 'smile',
          component: './Welcome',
        },
      ],
    },
    {
      name: '标准列表',
      icon: 'smile',
      path: '/listbasiclist',
      component: './ListBasicList',
    },
    {
      name: '卡片列表',
      icon: 'smile',
      path: '/listcardlist',
      component: './ListCardList',
    },
    {
      name: '搜索列表（应用）',
      icon: 'smile',
      path: '/listsearchapplications',
      component: './ListSearchApplications',
    },
    {
      name: '搜索列表',
      icon: 'smile',
      path: '/listsearch',
      component: './ListSearch',
    },
    {
      name: '查询表格',
      icon: 'smile',
      path: '/listtablelist',
      component: './ListTableListTwo',
    },
    {
      name: 'list.table-list',
      icon: 'table',
      path: '/list',
      component: './ListTableList',
    },
    {
      path: '/',
      redirect: '/welcome',
    },
    {
      name: '分析页',
      icon: 'smile',
      path: '/dashboardanalysis',
      component: './DashboardAnalysis',
    },
    {
      name: '个人中心',
      icon: 'smile',
      path: '/accountcenter',
      component: './AccountCenter',
    },
    {
      name: '个人设置',
      icon: 'smile',
      path: '/accountsettings',
      component: './AccountSettings',
    },
    {
      name: '监控页',
      icon: 'smile',
      path: '/dashboardmonitor',
      component: './DashboardMonitor',
    },
    {
      name: '工作台',
      icon: 'smile',
      path: '/dashboardworkplace',
      component: './DashboardWorkplace',
    },
    {
      name: '高级表单',
      icon: 'smile',
      path: '/formadvancedform',
      component: './FormAdvancedForm',
    },
    {
      name: '分步表单',
      icon: 'smile',
      path: '/formstepform',
      component: './FormStepForm',
    },
    {
      name: '高级详情页',
      icon: 'smile',
      path: '/profileadvanced',
      component: './ProfileAdvanced',
    },
    {
      name: '基础详情页',
      icon: 'smile',
      path: '/profilebasic',
      component: './ProfileBasic',
    },
    {
      name: '注册结果页',
      icon: 'smile',
      path: '/demo',
      component: './DemoFour',
    },
    {
      name: '成功页',
      icon: 'smile',
      path: '/demo',
      component: './DemoThree',
    },
    {
      name: '失败页',
      icon: 'smile',
      path: '/demo',
      component: './DemoTwo',
    },
    {
      name: '搜索列表（文章）',
      icon: 'smile',
      path: '/demo',
      component: './Demo',
    },
    {
      component: './404',
    },
  ],
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    // ...darkTheme,
    'primary-color': defaultSettings.primaryColor,
  },
  // @ts-ignore
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy[REACT_APP_ENV || 'dev'],
  manifest: {
    basePath: '/',
  },
});
