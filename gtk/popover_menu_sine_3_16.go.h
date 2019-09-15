// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_16

static GtkPopoverMenu *
toGtkPopoverMenu(void *p)
{
	return (GTK_POPOVER_MENU(p));
}