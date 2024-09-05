import { NbMenuItem } from '@nebular/theme';

export const MENU_ITEMS: NbMenuItem[] = [
  {
    title: 'Bảng điều khiển',
    icon: 'shopping-cart-outline',
    link: '/fund/dashboard',
    home: true,
  },
  {
    title: 'Lịch sử giao dịch',
    icon: 'shopping-cart-outline',
    link: '/fund/trade-history',
    home: false,
  },
  {
    title: 'Biểu đồ',
    icon: 'shopping-cart-outline',
    link: '/fund/e-chart',
    home: false,
  },
];
