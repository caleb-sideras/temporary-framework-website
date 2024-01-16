import { customElement, property } from 'lit/decorators.js';
import { TListItem2 } from "./t-list-item-2";
import { PropertyValueMap } from 'lit';

@customElement('t-dropdown-list-item')
export class TDropdownListItem extends TListItem2 {

  // @property({ type: Number, attribute: 'tabindex' }) tabIndex = -1;

  // protected updated(_changedProperties: PropertyValueMap<TDropdownListItem>): void {
  //   if (_changedProperties.has('tabIndex') && this.tabIndex) {
  //     this.dispatchEvent(this.createExternalActivationEvent());
  //   }
  // }
  static get observedAttributes() {
    return ["tabindex", ...super.observedAttributes];
  }

  attributeChangedCallback(name: string, oldValue: string, newValue: string) {

    super.attributeChangedCallback(name, oldValue, newValue);

    if (oldValue && name === "tabindex" && oldValue != newValue) {
      this.dispatchEvent(this.createExternalActivationEvent());
    }
  }

  createExternalActivationEvent() {
    return new Event('external-activation', { bubbles: true, composed: true });
  }
}