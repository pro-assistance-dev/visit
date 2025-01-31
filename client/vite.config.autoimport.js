import { pfName } from 'sprof';

export default {
  include: [
    '.ts',
    /\.[tj]sx?$/, // .ts, .tsx, .js, .jsx
    /\.vue$/,
    /\.vue\?vue/, // .vue
    /\.md$/, // .md
  ],
  imports: [
    'vue',
    'vitest',
    {
      from: 'sprof',
      imports: pfName,
    },
    {
      from: '@/store/index',
      imports: ['Store'],
    },
  ],
  types: true,
  defaultExportByFilenam: true,
  dirs: ['srs/modules/**', 'srs/services/**'],
  // ignoreDts: ['srs/services', 'Message'],
  vueTemplate: true,
  dts: true,
  eslintrc: {
    enabled: true,
  },
};
