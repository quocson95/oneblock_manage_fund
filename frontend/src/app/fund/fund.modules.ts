import { NgModule } from '@angular/core';
import { NbCardModule, NbIconModule, NbInputModule, NbListModule, NbMenuModule, NbTreeGridModule } from '@nebular/theme';

import { ThemeModule } from '../@theme/theme.module';
import { MiscellaneousModule } from './miscellaneous/miscellaneous.module';
import { FundRoutingModule } from './fund-routing.module';
import { FundDashboardModule } from './dashboard/dashboard.module';
import { FundComponent } from './fund.component';
import { TradeHistoryComponent } from './trade-history/trade-history.component';
import { Ng2SmartTableModule } from 'ng2-smart-table';
import { EChartComponent } from './echart/echart.component';


@NgModule({
    imports: [
        FundRoutingModule,
        ThemeModule,
        NbMenuModule,
        MiscellaneousModule,
        FundDashboardModule,
        Ng2SmartTableModule,
        NbTreeGridModule,
        NbIconModule,
        NbInputModule,
        ThemeModule,
        Ng2SmartTableModule,
        NbCardModule,
        NbListModule,
    ],
    declarations: [
        FundComponent,
        TradeHistoryComponent,
        EChartComponent,
    ],
})
export class FundModule {
}