import { Component } from '@angular/core';

@Component({
  selector: 'ngx-echart',
  templateUrl: 'echart.component.html',
  styleUrls: ['echart.component.scss'],
})
export class EChartComponent {

  links: { name: string, url: string }[] = [
    { name: 'SP500', url: 'https://api.oneblock.vn/be/chart/sp500.html' },
    { name: 'BTC ETH Statics', url: 'https://api.oneblock.vn/be/chart/btc_eth_stat.html' },
    { name: 'BTC Gold', url: 'https://api.oneblock.vn/be/chart/btc_gold.html' },
    { name: 'BTC Holder', url: 'https://api.oneblock.vn/be/chart/btc_holder.html' },
    { name: 'Funding Market Corelation', url: 'https://api.oneblock.vn/be/chart/funding_market_core.html' },
    { name: 'ETH Gas & BTC', url: 'https://api.oneblock.vn/be/chart/eth_gas_btc.html' },
    { name: 'VND.USD & BTC', url: 'https://api.oneblock.vn/be/chart/usd_vnd_btc.html' },

  ];
}
