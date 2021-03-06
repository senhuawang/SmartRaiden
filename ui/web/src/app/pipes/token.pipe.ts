import { Pipe, PipeTransform } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { SelectItem } from 'primeng/primeng';

import { SmartRaidenService } from '../services/smartraiden.service';
import { Usertoken } from '../models/usertoken';

@Pipe({
  name: 'token'
})
export class TokenPipe implements PipeTransform {

  constructor(private smartraidenService: SmartRaidenService) {}

  tokenToString(token: Usertoken): string {
    let text = '';
    if (!token) {
      return '';
    }
    if (token.symbol) {
      text += `[${token.symbol}] `;
    }
    if (token.name) {
      text += `${token.name} `;
    }
    if (text) {
      text += `(${token.address})`;
    } else {
      text = token.address;
    }
    return text;
  }

  tokensToSelectItems(tokens: Array<Usertoken>): Array<SelectItem> {
    return tokens.map((token) => ({
      value: token.address,
      label: this.tokenToString(token)
    }));
  }

  transform(address: string, args?: any): Observable<string> {
    return this.smartraidenService.getUsertoken(address, false)
      .map((token) => this.tokenToString(token));
  }

}
