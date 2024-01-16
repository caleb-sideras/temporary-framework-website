import { customElement, property } from 'lit/decorators.js';
import { html } from 'lit';
import { TListItem } from './t-list-item';
import { createRequestActivationEvent, } from '@material/web/list/internal/list-navigation-helpers.js';

@customElement('t-dropdown-title')
export class TDropdownTitle extends TListItem {

  @property({ type: Boolean, reflect: true }) open = false;

  protected override render() {
    return this.renderListItem(html`
      <md-item>
        <div slot="container">
          ${this.renderRipple()} ${this.renderFocusRing()}
        </div>
        <slot name="start" slot="start"></slot>
        <slot name="end" slot="end">  
          ${this.open ? this.inactiveSlot : this.activeSlot}
        </slot>
        ${this.renderBody()}
      </md-item>
    `);
  }

  get activeSlot() {
    return html`<slot name="active"></slot>`;
  }

  get inactiveSlot() {
    return html`<slot name="inactive"></slot>`;
  }

  protected override onClick() {
    console.log("CLICK")
    super.onClick();

    this.open = !this.open;
    this.dispatchEvent(new Event('title-activation', { bubbles: true, composed: true }));
  }
}
