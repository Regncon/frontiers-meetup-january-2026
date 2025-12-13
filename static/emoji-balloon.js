class EmojiBalloon extends HTMLElement {
	constructor() {
		super();
		this._storageKey = null;
		this._hasConnected = false;
		this._activeAnimation = null;
	}

	static get observedAttributes() {
		return ["data-emoji", "data-count"];
	}

	connectedCallback() {
		this._hasConnected = true;
		this._runCheck();
	}

	attributeChangedCallback(name, oldValue, newValue) {
		if (!this.isConnected) return;
		if (!this._hasConnected) return;
		if (oldValue === newValue) return;

		this._runCheck();
	}

	_runCheck() {
		const emoji = (this.getAttribute("data-emoji") || "").trim();

		const countStrRaw = this.getAttribute("data-count") || "0";
		const countStr = countStrRaw.trim();
		const count = Number(countStr);

		this._storageKey = "fmj26:emojiCount:" + encodeURIComponent(emoji);

		const storedRaw = sessionStorage.getItem(this._storageKey);
		const stored = storedRaw === null ? null : Number(storedRaw);

		if (!Number.isFinite(count)) {
			console.warn("[emoji-balloon] count is not a number:", countStrRaw);
			return;
		}

		// No animation on first load per emoji per tab.
		if (storedRaw === null) {
			sessionStorage.setItem(this._storageKey, String(count));
			return;
		}

		if (Number.isFinite(stored) && count > stored) {
			this._animate();
		}

		sessionStorage.setItem(this._storageKey, String(count));
	}

	_animate() {
		const target = this.querySelector("[data-emoji-balloon-target]");
		if (!target) return;

		if (this._activeAnimation) {
			this._activeAnimation.cancel();
			this._activeAnimation = null;
		}

		this._activeAnimation = target.animate(
			[
				{ transform: "translateY(0px) scale(1)" },
				{ transform: "translateY(-12px) scale(1.14)", offset: 0.45 },
				{ transform: "translateY(-18px) scale(1.02)" },
			],
			{
				duration: 320,
				easing: "cubic-bezier(0.2, 0.9, 0.2, 1)",
				fill: "none",
			}
		);
	}
}

if (!customElements.get("emoji-balloon")) {
	customElements.define("emoji-balloon", EmojiBalloon);
}
