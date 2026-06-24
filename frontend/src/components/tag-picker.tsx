"use client"

import { useEffect, useState } from "react"
import { api } from "@/lib/axios"
import { Tags, Loader2 } from "lucide-react"
import { TagBadge } from "./tag-badge"

import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

interface TagItem {
  id: string
  name: string
  color?: string | null
}

interface TagPickerProps {
  itemId: string
  itemType: "KANJI" | "VOCABULARY" | "GRAMMAR"
  attachedTags?: TagItem[]
  onTagsChange?: (tags: TagItem[]) => void
}

export function TagPicker({ itemId, itemType, attachedTags = [], onTagsChange }: TagPickerProps) {
  const [allTags, setAllTags] = useState<TagItem[]>([])
  const [currentTags, setCurrentTags] = useState<TagItem[]>(attachedTags)
  const [isToggling, setIsToggling] = useState<string | null>(null)

  useEffect(() => {
    const fetchTags = async () => {
      try {
        const [allRes, itemRes] = await Promise.all([
          api.get("/tags"),
          api.get(`/tags/item/${itemType}/${itemId}`),
        ])
        setAllTags(allRes.data.data || [])
        const itemTags = itemRes.data.data || []
        setCurrentTags(itemTags)
      } catch (err) {
        console.error("Failed to load tags", err)
      }
    }
    fetchTags()
  }, [itemId, itemType])

  const toggleTag = async (tag: TagItem) => {
    const isAttached = currentTags.some((t) => t.id === tag.id)
    setIsToggling(tag.id)
    try {
      if (isAttached) {
        await api.delete("/tags/attach", {
          data: { tagId: tag.id, itemType, itemId },
        })
        const updated = currentTags.filter((t) => t.id !== tag.id)
        setCurrentTags(updated)
        onTagsChange?.(updated)
      } else {
        await api.post("/tags/attach", {
          tagId: tag.id,
          itemType,
          itemId,
        })
        const updated = [...currentTags, tag]
        setCurrentTags(updated)
        onTagsChange?.(updated)
      }
    } catch (err) {
      console.error("Failed to toggle tag", err)
    } finally {
      setIsToggling(null)
    }
  }

  return (
    <div className="flex flex-col gap-1.5">
      <div className="flex flex-wrap items-center gap-1">
        {currentTags.map((tag) => (
          <TagBadge key={tag.id} tag={tag} onRemove={() => toggleTag(tag)} />
        ))}
      </div>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" size="sm" className="h-7 gap-1 text-xs text-muted-foreground">
            <Tags className="h-3 w-3" />
            {currentTags.length === 0 ? "Add Tag" : "Edit Tags"}
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="max-h-64 overflow-y-auto" align="start">
          <DropdownMenuLabel>Tags</DropdownMenuLabel>
          <DropdownMenuSeparator />
          {allTags.length === 0 ? (
            <div className="px-2 py-1.5 text-xs text-muted-foreground">
              No tags created yet
            </div>
          ) : (
            allTags.map((tag) => (
              <DropdownMenuCheckboxItem
                key={tag.id}
                checked={currentTags.some((t) => t.id === tag.id)}
                onCheckedChange={() => toggleTag(tag)}
                disabled={isToggling === tag.id}
              >
                {isToggling === tag.id ? (
                  <Loader2 className="h-3 w-3 animate-spin" />
                ) : (
                  <span
                    className="h-2.5 w-2.5 rounded-full shrink-0"
                    style={{ backgroundColor: tag.color || "#78716c" }}
                  />
                )}
                {tag.name}
              </DropdownMenuCheckboxItem>
            ))
          )}
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
