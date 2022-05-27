<script context="module" lang="ts">
    export enum BadgeType {
        INFO,
        ERROR,
    }

    function getFontClass(type: BadgeType) {
        switch (type) {
            case BadgeType.INFO:
                return "text-gray-200 font-light"
            case BadgeType.ERROR:
                return "text-red-500 font-bold"
        }
    }

    function getIconClass(type: BadgeType) {
        switch (type) {
            case BadgeType.INFO:
                return "stroke-gray-200"
            case BadgeType.ERROR:
                return "stroke-red-500"
        }
    }
</script>

<script lang="ts">
    export let type: BadgeType = BadgeType.INFO
    export let label: string = ""
</script>

<span
    class="flex text-xs w-36 overflow-hidden px-2 gap-1 items-center {getFontClass(
        type
    )}"
    data-cy="badge"
>
    <span class="stroke-gray-200 h-3 w-3 {getIconClass(type)}">
        <slot />
    </span>
    {label}
</span>
