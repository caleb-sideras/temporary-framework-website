// as of now, htmx does not support shadow dom
// merge https://github.com/bigskysoftware/htmx/pull/2075 !!!

import { LitElement, html } from 'lit';
import { property } from 'lit/decorators.js';

export class HTMXElement extends LitElement {

  @property({ type: String }) href = '';
  @property({ type: String, attribute: 'hx-boost' }) hxBoost = '';
  @property({ type: String, attribute: 'hx-get' }) hxGet = '';
  @property({ type: String, attribute: 'hx-post' }) hxPost = '';
  @property({ type: String, attribute: 'hx-patch' }) hxPatch = '';
  @property({ type: String, attribute: 'hx-put' }) hxPut = '';
  @property({ type: String, attribute: 'hx-delete' }) hxDelete = '';
  @property({ type: String, attribute: 'hx-trigger' }) hxTrigger = '';
  @property({ type: String, attribute: 'hx-indicator' }) hxIndicator = '';
  @property({ type: String, attribute: 'hx-target' }) hxTarget = '';
  @property({ type: String, attribute: 'hx-swap' }) hxSwap = '';
  @property({ type: String, attribute: 'hx-history-elt' }) hxHistoryElt = '';

  renderAnchor(content: unknown) {
    return html`
      <a
        href="${this.href}"
        hx-boost="${this.hxBoost}"
      >
        ${content}
      </a>
    `
  }

  renderDiv(content: unknown) {
    return html`
      <div
        hx-get="${this.hxGet}"
        hx-post="${this.hxPost}"
        hx-patch="${this.hxPatch}"
        hx-put="${this.hxPut}"
        hx-delete="${this.hxDelete}"
        hx-trigger="${this.hxTrigger}"
        hx-indicator="${this.hxIndicator}"
        hx-target="${this.hxTarget}"
        hx-swap="${this.hxSwap}"
        hx-history-elt="${this.hxHistoryElt}"      >
        ${content}
      </div>
    `
  }
}