import { RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';

import { NotFoundComponent } from './miscellaneous/not-found/not-found.component';
import { FundComponent } from './fund.component';
import { FundDashboardComponent } from './dashboard/dashboard.component';
import { TradeHistoryComponent } from './trade-history/trade-history.component';
import { EChartComponent as EChartComponent } from './echart/echart.component';

const routes: Routes = [{
    path: '',
    component: FundComponent,
    children: [
        {
            path: 'dashboard',
            component: FundDashboardComponent,
        },
        {
            path: 'trade-history',
            component: TradeHistoryComponent,
        },
        {
            path: 'e-chart',
            component: EChartComponent,
        },
        {
            path: '',
            redirectTo: 'dashboard',
            pathMatch: 'full',
        },
        {
            path: '**',
            component: NotFoundComponent,
        },

    ]
}];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class FundRoutingModule {
}