<script context="module" lang="ts">
    export enum ButtonType {
        PRIMARY,
        SECONDARY,
    }

    function getButtonClass(type: ButtonType) {
        switch (type) {
            case ButtonType.PRIMARY:
                return "bg-blue-600 hover:bg-blue-500 text-white"
            case ButtonType.SECONDARY:
                return "bg-gray-300 hover:bg-gray-200 text-slate-700"
        }
    }
</script>

<script lang="ts">
    export let label: string
    export let onClick: any = undefined
    export let fullWidth: boolean = false
    export let type: ButtonType = ButtonType.PRIMARY

    const buttonClass = getButtonClass(type)
</script>

<button
    on:click={onClick}
    class="drop-shadow-card font-bold rounded-xl py-2 px-5 {buttonClass}"
    class:w-full={fullWidth}
>
    <span>{label}</span>
</button>
