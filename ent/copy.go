// Code generated by ent, DO NOT EDIT.

package ent

// CopyRepository allows to create a new Repository copying the existing
// values of input.
func (rc *RepositoryCreate) CopyRepository(input *Repository) *RepositoryCreate {
	rc.SetID(input.ID)
	rc.SetNodeID(input.NodeID)
	rc.SetName(input.Name)
	rc.SetFullName(input.FullName)
	rc.SetPrivate(input.Private)
	rc.SetHTMLURL(input.HTMLURL)
	rc.SetDescription(input.Description)
	rc.SetFork(input.Fork)
	rc.SetURL(input.URL)
	rc.SetArchiveURL(input.ArchiveURL)
	rc.SetAssigneesURL(input.AssigneesURL)
	rc.SetBlobsURL(input.BlobsURL)
	rc.SetBranchesURL(input.BranchesURL)
	rc.SetCollaboratorsURL(input.CollaboratorsURL)
	rc.SetCommentsURL(input.CommentsURL)
	rc.SetCommitsURL(input.CommitsURL)
	rc.SetCompareURL(input.CompareURL)
	rc.SetContentsURL(input.ContentsURL)
	rc.SetContributorsURL(input.ContributorsURL)
	rc.SetDeploymentsURL(input.DeploymentsURL)
	rc.SetDownloadsURL(input.DownloadsURL)
	rc.SetEventsURL(input.EventsURL)
	rc.SetForksURL(input.ForksURL)
	rc.SetGitCommitsURL(input.GitCommitsURL)
	rc.SetGitRefsURL(input.GitRefsURL)
	rc.SetGitTagsURL(input.GitTagsURL)
	rc.SetGitURL(input.GitURL)
	rc.SetIssueCommentURL(input.IssueCommentURL)
	rc.SetIssueEventsURL(input.IssueEventsURL)
	rc.SetIssuesURL(input.IssuesURL)
	rc.SetKeysURL(input.KeysURL)
	rc.SetLabelsURL(input.LabelsURL)
	rc.SetLanguagesURL(input.LanguagesURL)
	rc.SetMergesURL(input.MergesURL)
	rc.SetMilestonesURL(input.MilestonesURL)
	rc.SetNotificationsURL(input.NotificationsURL)
	rc.SetPullsURL(input.PullsURL)
	rc.SetReleasesURL(input.ReleasesURL)
	rc.SetSSHURL(input.SSHURL)
	rc.SetStargazersURL(input.StargazersURL)
	rc.SetStatusesURL(input.StatusesURL)
	rc.SetSubscribersURL(input.SubscribersURL)
	rc.SetSubscriptionURL(input.SubscriptionURL)
	rc.SetTagsURL(input.TagsURL)
	rc.SetTeamsURL(input.TeamsURL)
	rc.SetTreesURL(input.TreesURL)
	rc.SetCloneURL(input.CloneURL)
	rc.SetMirrorURL(input.MirrorURL)
	rc.SetHooksURL(input.HooksURL)
	rc.SetSvnURL(input.SvnURL)
	rc.SetHomepage(input.Homepage)
	rc.SetLanguage(input.Language)
	rc.SetForksCount(input.ForksCount)
	rc.SetStargazersCount(input.StargazersCount)
	rc.SetWatchersCount(input.WatchersCount)
	rc.SetSize(input.Size)
	rc.SetDefaultBranch(input.DefaultBranch)
	rc.SetOpenIssuesCount(input.OpenIssuesCount)
	rc.SetIsTemplate(input.IsTemplate)
	rc.SetTopics(input.Topics)
	rc.SetHasIssues(input.HasIssues)
	rc.SetHasProjects(input.HasProjects)
	rc.SetHasWiki(input.HasWiki)
	rc.SetHasPages(input.HasPages)
	rc.SetHasDownloads(input.HasDownloads)
	rc.SetHasDiscussions(input.HasDiscussions)
	rc.SetArchived(input.Archived)
	rc.SetDisabled(input.Disabled)
	rc.SetVisibility(input.Visibility)
	rc.SetPushedAt(input.PushedAt)
	rc.SetCreatedAt(input.CreatedAt)
	rc.SetUpdatedAt(input.UpdatedAt)
	rc.SetSubscribersCount(input.SubscribersCount)
	rc.SetNetworkCount(input.NetworkCount)
	rc.SetForks(input.Forks)
	rc.SetOpenIssues(input.OpenIssues)
	rc.SetWatchers(input.Watchers)
	return rc
}

// CopyRepository allows to update a Repository copying the existing
// values of input.
func (ruo *RepositoryUpdateOne) CopyRepository(input *Repository) *RepositoryUpdateOne {
	ruo.SetNodeID(input.NodeID)
	ruo.SetName(input.Name)
	ruo.SetFullName(input.FullName)
	ruo.SetPrivate(input.Private)
	ruo.SetHTMLURL(input.HTMLURL)
	ruo.SetDescription(input.Description)
	ruo.SetFork(input.Fork)
	ruo.SetURL(input.URL)
	ruo.SetArchiveURL(input.ArchiveURL)
	ruo.SetAssigneesURL(input.AssigneesURL)
	ruo.SetBlobsURL(input.BlobsURL)
	ruo.SetBranchesURL(input.BranchesURL)
	ruo.SetCollaboratorsURL(input.CollaboratorsURL)
	ruo.SetCommentsURL(input.CommentsURL)
	ruo.SetCommitsURL(input.CommitsURL)
	ruo.SetCompareURL(input.CompareURL)
	ruo.SetContentsURL(input.ContentsURL)
	ruo.SetContributorsURL(input.ContributorsURL)
	ruo.SetDeploymentsURL(input.DeploymentsURL)
	ruo.SetDownloadsURL(input.DownloadsURL)
	ruo.SetEventsURL(input.EventsURL)
	ruo.SetForksURL(input.ForksURL)
	ruo.SetGitCommitsURL(input.GitCommitsURL)
	ruo.SetGitRefsURL(input.GitRefsURL)
	ruo.SetGitTagsURL(input.GitTagsURL)
	ruo.SetGitURL(input.GitURL)
	ruo.SetIssueCommentURL(input.IssueCommentURL)
	ruo.SetIssueEventsURL(input.IssueEventsURL)
	ruo.SetIssuesURL(input.IssuesURL)
	ruo.SetKeysURL(input.KeysURL)
	ruo.SetLabelsURL(input.LabelsURL)
	ruo.SetLanguagesURL(input.LanguagesURL)
	ruo.SetMergesURL(input.MergesURL)
	ruo.SetMilestonesURL(input.MilestonesURL)
	ruo.SetNotificationsURL(input.NotificationsURL)
	ruo.SetPullsURL(input.PullsURL)
	ruo.SetReleasesURL(input.ReleasesURL)
	ruo.SetSSHURL(input.SSHURL)
	ruo.SetStargazersURL(input.StargazersURL)
	ruo.SetStatusesURL(input.StatusesURL)
	ruo.SetSubscribersURL(input.SubscribersURL)
	ruo.SetSubscriptionURL(input.SubscriptionURL)
	ruo.SetTagsURL(input.TagsURL)
	ruo.SetTeamsURL(input.TeamsURL)
	ruo.SetTreesURL(input.TreesURL)
	ruo.SetCloneURL(input.CloneURL)
	ruo.SetMirrorURL(input.MirrorURL)
	ruo.SetHooksURL(input.HooksURL)
	ruo.SetSvnURL(input.SvnURL)
	ruo.SetHomepage(input.Homepage)
	ruo.SetLanguage(input.Language)
	ruo.SetForksCount(input.ForksCount)
	ruo.SetStargazersCount(input.StargazersCount)
	ruo.SetWatchersCount(input.WatchersCount)
	ruo.SetSize(input.Size)
	ruo.SetDefaultBranch(input.DefaultBranch)
	ruo.SetOpenIssuesCount(input.OpenIssuesCount)
	ruo.SetIsTemplate(input.IsTemplate)
	ruo.SetTopics(input.Topics)
	ruo.SetHasIssues(input.HasIssues)
	ruo.SetHasProjects(input.HasProjects)
	ruo.SetHasWiki(input.HasWiki)
	ruo.SetHasPages(input.HasPages)
	ruo.SetHasDownloads(input.HasDownloads)
	ruo.SetHasDiscussions(input.HasDiscussions)
	ruo.SetArchived(input.Archived)
	ruo.SetDisabled(input.Disabled)
	ruo.SetVisibility(input.Visibility)
	ruo.SetPushedAt(input.PushedAt)
	ruo.SetCreatedAt(input.CreatedAt)
	ruo.SetUpdatedAt(input.UpdatedAt)
	ruo.SetSubscribersCount(input.SubscribersCount)
	ruo.SetNetworkCount(input.NetworkCount)
	ruo.SetForks(input.Forks)
	ruo.SetOpenIssues(input.OpenIssues)
	ruo.SetWatchers(input.Watchers)
	return ruo
}
