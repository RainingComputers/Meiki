<script context="module" lang="ts">
    export enum ButtonType {
        HERO,
        PRIMARY,
        SECONDARY,
    }

    function getButtonClass(type: ButtonType) {
        switch (type) {
            case ButtonType.HERO:
                return "bg-background-primary hover:bg-background-primary-focus text-content-primary text-2xl py-4 px-10"
            case ButtonType.PRIMARY:
                return "bg-background-primary hover:bg-background-primary-focus text-content-primary py-2 px-5"
            case ButtonType.SECONDARY:
                return "bg-background-secondary hover:bg-background-secondary-focus text-content-secondary py-2 px-5"
        }
    }
</script>

<script lang="ts">
    export let label: string
    export let onClick: () => void = undefined
    export let fullWidth: boolean = false
    export let expand: boolean = false
    export let type: ButtonType = ButtonType.PRIMARY

    const buttonClass = getButtonClass(type)
</script>

<button
    on:click={onClick}
    class="drop-shadow-card font-bold rounded-xl {buttonClass}"
    class:w-full={fullWidth}
    class:flex-grow={expand}
>
    <span>{label}</span>
</button>
