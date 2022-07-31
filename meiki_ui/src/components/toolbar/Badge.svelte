<script context="module" lang="ts">
    export enum BadgeType {
        INFO,
        ERROR,
    }

    function getFontClass(type: BadgeType) {
        switch (type) {
            case BadgeType.INFO:
                return "text-toolbar-content font-light"
            case BadgeType.ERROR:
                return "text-error-content font-bold"
        }
    }

    function getIconClass(type: BadgeType) {
        switch (type) {
            case BadgeType.INFO:
                return "stroke-content-title"
            case BadgeType.ERROR:
                return "stroke-content-error"
        }
    }
</script>

<script lang="ts">
    export let type: BadgeType = BadgeType.INFO
    export let label: string = ""
</script>

<span
    class="flex text-xs w-36 overflow-hidden px-2 gap-1 items-center {getFontClass(type)}"
    data-cy="badge"
>
    <span class="stroke-toolbar-content h-3 w-3 {getIconClass(type)}">
        <slot />
    </span>
    {label}
</span>
