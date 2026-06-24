"use client"

import { X } from "lucide-react"

interface TagBadgeProps {
  tag: { id: string; name: string; color?: string | null }
  onRemove?: () => void
}

export function TagBadge({ tag, onRemove }: TagBadgeProps) {
  return (
    <span
      className="inline-flex items-center gap-1 rounded-full px-2 py-0.5 text-xs font-medium text-white"
      style={{ backgroundColor: tag.color || "#78716c" }}
    >
      {tag.name}
      {onRemove && (
        <button
          type="button"
          onClick={(e) => {
            e.stopPropagation()
            onRemove()
          }}
          className="ml-0.5 rounded-full hover:bg-white/20 p-0.5 transition-colors"
        >
          <X className="h-2.5 w-2.5" />
        </button>
      )}
    </span>
  )
}
