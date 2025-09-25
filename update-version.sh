#!/bin/bash

# Version update script for nav
# Usage: ./update-version.sh [new_version]

set -e

# Check if version argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <new_version>"
    echo "Example: $0 1.0.3"
    exit 1
fi

NEW_VERSION=$1

# Validate version format (basic semver check)
if ! [[ $NEW_VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in semver format (e.g., 1.0.3)"
    exit 1
fi

echo "Updating version to $NEW_VERSION..."

# Update the version file
echo "$NEW_VERSION" > version

# Test the build to make sure everything still works
echo "Testing build..."
go build -o nav

# Test the version command
VERSION_OUTPUT=$(./nav version)
if [[ $VERSION_OUTPUT == *"$NEW_VERSION"* ]]; then
    echo "✓ Version update successful: $VERSION_OUTPUT"
else
    echo "✗ Version update failed. Expected $NEW_VERSION, got: $VERSION_OUTPUT"
    exit 1
fi

# Clean up test binary
rm -f nav

# Stage changes
echo "Staging changes..."
git add version

# Check if there are any other changes that need to be committed
if ! git diff --cached --quiet; then
    # Commit the version update
    echo "Committing version update..."
    git commit -m "bump version to v$NEW_VERSION"
    
    # Create a git tag
    echo "Creating git tag v$NEW_VERSION..."
    git tag "v$NEW_VERSION"
    
    # Push changes and tags
    echo "Pushing to GitHub..."
    git push origin HEAD
    git push origin "v$NEW_VERSION"
    
    echo "✓ Version $NEW_VERSION successfully updated and pushed to GitHub!"
    echo "✓ Git tag v$NEW_VERSION created and pushed"
else
    echo "No changes to commit. Version file may already be at $NEW_VERSION"
fi

echo "Done!"
